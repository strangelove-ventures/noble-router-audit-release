package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble/x/router/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q QueryServer) InFlightPacket(c context.Context, req *types.QueryGetInFlightPacketRequest) (*types.QueryGetInFlightPacketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := q.keeper.GetInFlightPacket(ctx, req.ChannelId, req.PortId, req.Sequence)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInFlightPacketResponse{InFlightPacket: val}, nil
}

func (q QueryServer) InFlightPackets(c context.Context, req *types.QueryAllInFlightPacketsRequest) (*types.QueryAllInFlightPacketsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	inFlightPackets, pageRes, err := q.keeper.GetAllInFlightPacketsPaginated(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInFlightPacketsResponse{InFlightPackets: inFlightPackets, Pagination: pageRes}, nil
}
