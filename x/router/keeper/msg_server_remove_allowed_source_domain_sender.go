package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/strangelove-ventures/noble/x/router/types"
)

func (m msgServer) RemoveAllowedSourceDomainSender(goCtx context.Context, msg *types.MsgRemoveAllowedSourceDomainSender) (*types.MsgRemoveAllowedSourceDomainSenderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner := m.keeper.GetOwner(ctx)
	if owner != msg.From {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "this message sender cannot remove allowed source domain senders")
	}

	allowed := m.keeper.IsAllowedSourceDomainSender(ctx, msg.DomainId, msg.Address)
	if !allowed {
		return nil, types.ErrAllowedSourceDomainSenderNotFound
	}

	m.keeper.DeleteAllowedSourceDomainSender(ctx, msg.DomainId, msg.Address)

	event := types.AllowedSourceDomainSenderRemoved{
		Domain:  msg.DomainId,
		Address: msg.Address,
	}
	err := ctx.EventManager().EmitTypedEvent(&event)

	return &types.MsgRemoveAllowedSourceDomainSenderResponse{}, err
}
