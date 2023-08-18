package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble/x/cctp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Roles(goCtx context.Context, req *types.QueryRolesRequest) (*types.QueryRolesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	roles := types.QueryRolesResponse{
		Owner:           k.GetOwner(ctx),
		AttesterManager: k.GetAttesterManager(ctx),
		Pauser:          k.GetPauser(ctx),
		TokenController: k.GetTokenController(ctx),
	}

	return &roles, nil
}
