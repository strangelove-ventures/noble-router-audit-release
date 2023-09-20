package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble/x/router/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q QueryServer) Mint(c context.Context, req *types.QueryGetMintRequest) (*types.QueryGetMintResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := q.keeper.GetMint(ctx, req.SourceDomain, req.Nonce)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMintResponse{Mint: val}, nil
}

func (q QueryServer) Mints(c context.Context, req *types.QueryAllMintsRequest) (*types.QueryAllMintsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	mints, pageRes, err := q.keeper.GetAllMintsPaginated(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMintsResponse{Mints: mints, Pagination: pageRes}, nil
}
