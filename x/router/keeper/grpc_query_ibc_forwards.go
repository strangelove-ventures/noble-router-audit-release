package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble/x/router/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q QueryServer) IBCForward(c context.Context, req *types.QueryGetIBCForwardRequest) (*types.QueryGetIBCForwardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := q.keeper.GetIBCForward(ctx, req.SourceDomain, req.Nonce)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetIBCForwardResponse{IbcForward: val}, nil
}

func (q QueryServer) IBCForwards(c context.Context, req *types.QueryAllIBCForwardsRequest) (*types.QueryAllIBCForwardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	ibcForwards, pageRes, err := q.keeper.GetAllIBCForwardsPaginated(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIBCForwardsResponse{IbcForwards: ibcForwards, Pagination: pageRes}, nil
}
