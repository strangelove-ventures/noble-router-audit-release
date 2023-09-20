package keeper

import (
	"context"

	"github.com/strangelove-ventures/noble/x/router/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) AcceptOwner(goCtx context.Context, msg *types.MsgAcceptOwner) (*types.MsgAcceptOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentOwner := m.keeper.GetOwner(ctx)
	pendingOwner, found := m.keeper.GetPendingOwner(ctx)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "pending owner is not set")
	}

	if pendingOwner != msg.From {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "you are not the pending owner")
	}

	m.keeper.SetOwner(ctx, pendingOwner)
	m.keeper.DeletePendingOwner(ctx)

	event := types.OwnerUpdated{
		PreviousOwner: currentOwner,
		NewOwner:      pendingOwner,
	}
	err := ctx.EventManager().EmitTypedEvent(&event)

	return &types.MsgAcceptOwnerResponse{}, err
}
