package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble/x/cctp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PerMessageBurnLimit(c context.Context, req *types.QueryGetPerMessageBurnLimitRequest) (*types.QueryGetPerMessageBurnLimitResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPerMessageBurnLimit(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPerMessageBurnLimitResponse{BurnLimit: val}, nil
}

func (k Keeper) PerMessageBurnLimits(c context.Context, req *types.QueryAllPerMessageBurnLimitsRequest) (*types.QueryAllPerMessageBurnLimitsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var perMessageBurnLimits []types.PerMessageBurnLimit
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	perMessageBurnLimitsStore := prefix.NewStore(store, types.KeyPrefix(types.PerMessageBurnLimitKeyPrefix))

	pageRes, err := query.Paginate(perMessageBurnLimitsStore, req.Pagination, func(key []byte, value []byte) error {
		var perMessageBurnLimit types.PerMessageBurnLimit
		if err := k.cdc.Unmarshal(value, &perMessageBurnLimit); err != nil {
			return err
		}

		perMessageBurnLimits = append(perMessageBurnLimits, perMessageBurnLimit)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPerMessageBurnLimitsResponse{BurnLimits: perMessageBurnLimits, Pagination: pageRes}, nil
}
