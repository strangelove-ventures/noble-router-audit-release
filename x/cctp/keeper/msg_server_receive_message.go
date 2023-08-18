package keeper

import (
	"bytes"
	"context"
	"strings"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/strangelove-ventures/noble/x/cctp/types"
	fiattokenfactorytypes "github.com/strangelove-ventures/noble/x/fiattokenfactory/types"
)

var (
	tokenMessengerRecipient = crypto.Keccak256([]byte("cctp/TokenMessenger"))

	zeroByteArray = []byte{ // 32 bytes
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}
)

func (k msgServer) ReceiveMessage(goCtx context.Context, msg *types.MsgReceiveMessage) (*types.MsgReceiveMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	paused, found := k.GetSendingAndReceivingMessagesPaused(ctx)
	if found && paused.Paused {
		return nil, sdkerrors.Wrap(types.ErrReceiveMessage, "sending and receiving messages are paused")
	}

	// Validate each signature in the attestation
	publicKeys := k.GetAllAttesters(ctx)
	if len(publicKeys) == 0 {
		return nil, sdkerrors.Wrap(types.ErrReceiveMessage, "no attesters found")
	}

	signatureThreshold, found := k.GetSignatureThreshold(ctx)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrReceiveMessage, "signature threshold not found")
	}

	if err := VerifyAttestationSignatures(msg.Message, msg.Attestation, publicKeys, signatureThreshold.Amount); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrReceiveMessage, "unable to verify signatures")
	}

	// parse message
	message, err := new(types.Message).Parse(msg.Message)
	if err != nil {
		return nil, err
	}

	// validate domain
	if message.DestinationDomain != types.NobleDomainId {
		return nil, sdkerrors.Wrapf(types.ErrReceiveMessage, "incorrect destination domain: %d", message.DestinationDomain)
	}

	// validate destination caller
	if !bytes.Equal(message.DestinationCaller, zeroByteArray) {
		bech32Prefix := sdk.GetConfig().GetBech32AccountAddrPrefix()
		destinationCaller, err := bech32.ConvertAndEncode(bech32Prefix, message.DestinationCaller[12:])
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrReceiveMessage, "unable to encode destination caller: %s", msg.From)
		}

		if destinationCaller != msg.From {
			return nil, sdkerrors.Wrapf(types.ErrReceiveMessage, "incorrect destination caller: %s, sender: %s", destinationCaller, msg.From)
		}
	}

	// validate version
	if message.Version != types.NobleMessageVersion {
		return nil, sdkerrors.Wrapf(types.ErrReceiveMessage, "incorrect message version. expected: %d, found: %d", types.NobleMessageVersion, message.Version)
	}

	// validate nonce is available
	// note: we use the domain/nonce combo instead of a hash
	usedNonce := types.Nonce{SourceDomain: message.SourceDomain, Nonce: message.Nonce}
	found = k.GetUsedNonce(ctx, usedNonce)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrReceiveMessage, "nonce already used")
	}

	// mark nonce as used
	k.SetUsedNonce(ctx, usedNonce)

	// verify and parse BurnMessage
	if bytes.Equal(message.Recipient, tokenMessengerRecipient) { // then mint
		burnMessage, err := new(types.BurnMessage).Parse(message.MessageBody)
		if err != nil {
			return nil, err
		}

		if burnMessage.Version != types.MessageBodyVersion {
			return nil, sdkerrors.Wrap(types.ErrReceiveMessage, "invalid message body version")
		}

		// look up Noble mint token from corresponding source domain/token
		tokenPair, found := k.GetTokenPair(ctx, message.SourceDomain, burnMessage.BurnToken)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrReceiveMessage, "corresponding noble mint token not found")
		}

		// get mint recipient as noble address
		bech32Prefix := sdk.GetConfig().GetBech32AccountAddrPrefix()
		mintRecipient, err := sdk.Bech32ifyAddressBytes(bech32Prefix, burnMessage.MintRecipient[12:])
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrReceiveMessage, "error bech32 encoding mint recipient address")
		}

		msgMint := fiattokenfactorytypes.MsgMint{
			From:    authtypes.NewModuleAddress(types.ModuleName).String(),
			Address: mintRecipient,
			Amount: sdk.Coin{
				Denom:  strings.ToLower(tokenPair.LocalToken),
				Amount: sdk.NewIntFromBigInt(burnMessage.Amount.BigInt()),
			},
		}
		_, err = k.fiattokenfactory.Mint(ctx, &msgMint)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "Error during minting")
		}

		mintEvent := types.MintAndWithdraw{
			MintRecipient: string(burnMessage.MintRecipient),
			Amount:        burnMessage.Amount,
			MintToken:     strings.ToLower(tokenPair.LocalToken),
		}
		err = ctx.EventManager().EmitTypedEvent(&mintEvent)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "Error emitting mint event")
		}
	}

	// on failure to decode, nil err from handleMessage
	if err := k.router.HandleMessage(ctx, msg.Message); err != nil {
		return nil, sdkerrors.Wrap(types.ErrHandleMessage, "Error in handleMessage")
	}

	event := types.MessageReceived{
		Caller:       msg.From,
		SourceDomain: message.SourceDomain,
		Nonce:        message.Nonce,
		Sender:       message.Sender,
		MessageBody:  message.MessageBody,
	}
	err = ctx.EventManager().EmitTypedEvent(&event)

	return &types.MsgReceiveMessageResponse{Success: true}, err
}
