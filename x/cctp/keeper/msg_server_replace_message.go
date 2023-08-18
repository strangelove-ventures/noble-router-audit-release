package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/noble/x/cctp/types"
)

func (k msgServer) ReplaceMessage(goCtx context.Context, msg *types.MsgReplaceMessage) (*types.MsgReplaceMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	paused, found := k.GetSendingAndReceivingMessagesPaused(ctx)
	if found && paused.Paused {
		return nil, sdkerrors.Wrap(types.ErrReplaceMessage, "sending and receiving messages are paused")
	}

	// Validate each signature in the attestation
	// Note: changing attesters or the signature threshold can render all previous messages irreplaceable
	attesters := k.GetAllAttesters(ctx)
	signatureThreshold, found := k.GetSignatureThreshold(ctx)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrReplaceMessage, "signature threshold not found")
	}

	if err := VerifyAttestationSignatures(msg.OriginalMessage, msg.OriginalAttestation, attesters, signatureThreshold.Amount); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrSignatureVerification, "unable to verify signatures")
	}

	// validate message format
	originalMessage, err := new(types.Message).Parse(msg.OriginalMessage)
	if err != nil {
		return nil, err
	}

	// validate that the original message sender is the same as this message sender
	if msg.From != string(originalMessage.Sender) {
		return nil, sdkerrors.Wrap(types.ErrReplaceMessage, "sender not permitted to use nonce")
	}

	// validate source domain
	if originalMessage.SourceDomain != types.NobleDomainId {
		return nil, sdkerrors.Wrap(types.ErrReplaceMessage, "message not originally sent from this domain")
	}

	err = k.sendMessage(
		ctx,
		originalMessage.DestinationDomain,
		originalMessage.Recipient,
		msg.NewDestinationCaller,
		originalMessage.Sender,
		originalMessage.Nonce,
		msg.NewMessageBody)

	return &types.MsgReplaceMessageResponse{}, err
}
