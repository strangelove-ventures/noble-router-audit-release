package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/strangelove-ventures/noble/x/cctp/types"
)

func (k msgServer) UnlinkTokenPair(goCtx context.Context, msg *types.MsgUnlinkTokenPair) (*types.MsgUnlinkTokenPairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenController := k.GetTokenController(ctx)
	if tokenController != msg.From {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "this message sender cannot unlink token pairs")
	}

	tokenPair, found := k.GetTokenPairHex(ctx, msg.RemoteDomain, msg.RemoteToken)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrTokenPairNotFound, "token pair doesn't exist in store")
	}

	k.DeleteTokenPair(ctx, msg.RemoteDomain, tokenPair.RemoteToken)

	event := types.TokenPairUnlinked{
		LocalToken:   tokenPair.LocalToken,
		RemoteDomain: tokenPair.RemoteDomain,
		RemoteToken:  msg.RemoteToken,
	}
	err := ctx.EventManager().EmitTypedEvent(&event)
	return &types.MsgUnlinkTokenPairResponse{}, err
}
