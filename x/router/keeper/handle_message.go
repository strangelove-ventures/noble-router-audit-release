package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/strangelove-ventures/noble-router/x/router/types"
	"strconv"
)

func (k Keeper) HandleMessage(goCtx context.Context, msg []byte) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// parse outer message
	outerMessage, err := decodeMessage(msg)
	if err != nil {
		return err
	}

	// parse internal message into IBCForward
	if ibcForward, err := decodeIBCForward(outerMessage.MessageBody); err == nil {
		if storedForward, ok := k.GetIBCForward(ctx, string(outerMessage.Sender), outerMessage.Nonce); ok {
			if storedForward.AckError {
				if existingMint, ok := k.GetMint(ctx, string(outerMessage.Sender), outerMessage.Nonce); ok {
					return k.ForwardPacket(ctx, ibcForward, existingMint)
				}
				panic("unexpected state")
			}

			return sdkerrors.Wrapf(types.ErrHandleMessage, "previous operation still in progress")
		}
		// this is the first time we are seeing this forward info, store it.
		k.SetIBCForward(ctx, types.StoreIBCForwardMetadata{
			SourceDomainSender: string(outerMessage.Sender),
			Nonce:              outerMessage.Nonce,
			Metadata:           &ibcForward,
		})
		if existingMint, ok := k.GetMint(ctx, string(outerMessage.Sender), outerMessage.Nonce); ok {
			return k.ForwardPacket(ctx, ibcForward, existingMint)
		}
		return nil
	}

	// try to parse internal message into burn (representing a remote burn -> local mint)
	if burnMessage, err := decodeBurnMessage(outerMessage.MessageBody); err == nil {
		// message is a Mint
		mint := types.Mint{
			SourceDomainSender: string(outerMessage.Sender),
			Nonce:              outerMessage.Nonce,
			Amount: &sdk.Coin{
				Denom:  string(burnMessage.BurnToken),
				Amount: sdk.NewInt(int64(burnMessage.Amount)),
			},
			DestinationDomain: strconv.Itoa(int(outerMessage.DestinationDomain)),
			MintRecipient:     string(burnMessage.MintRecipient),
		}
		k.SetMint(ctx, mint)
		if existingIBCForward, found := k.GetIBCForward(ctx, string(burnMessage.MessageSender), outerMessage.Nonce); found {
			return k.ForwardPacket(ctx, *existingIBCForward.Metadata, mint)
		}
	}

	return nil
}

func (k Keeper) ForwardPacket(ctx sdk.Context, ibcForward types.IBCForwardMetadata, mint types.Mint) error {
	timeout := ibcForward.TimeoutInNanoseconds
	if timeout == 0 {
		timeout = transfertypes.DefaultRelativePacketTimeoutTimestamp
	}

	transfer := &transfertypes.MsgTransfer{
		SourcePort:    ibcForward.Port,
		SourceChannel: ibcForward.Channel,
		Token:         *mint.Amount,
		Sender:        mint.MintRecipient,
		Receiver:      ibcForward.DestinationReceiver,
		TimeoutHeight: clienttypes.Height{
			RevisionNumber: 0,
			RevisionHeight: 0,
		},
		TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano()) + timeout,
		Memo:             ibcForward.Memo,
	}

	res, err := k.transferKeeper.Transfer(sdk.WrapSDKContext(ctx), transfer)
	if err != nil {
		return err
	}

	inFlightPacket := types.InFlightPacket{
		SourceDomainSender: mint.SourceDomainSender,
		Nonce:              mint.Nonce,
	}

	k.SetInFlightPacket(ctx, ibcForward.Channel, ibcForward.Port, res.Sequence, inFlightPacket)

	return nil
}
