package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble/x/router/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q QueryServer) AllowedSourceDomainSender(c context.Context, req *types.QueryAllowedSourceDomainSenderRequest) (*types.QueryAllowedSourceDomainSenderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	allowed := q.keeper.IsAllowedSourceDomainSender(ctx, req.DomainId, req.Address)
	if !allowed {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryAllowedSourceDomainSenderResponse{AllowedSourceDomainSender: types.AllowedSourceDomainSender{
		DomainId: req.DomainId,
		Address:  req.Address,
	}}, nil
}

func (q QueryServer) AllowedSourceDomainSenders(c context.Context, req *types.QueryAllowedSourceDomainSendersRequest) (*types.QueryAllowedSourceDomainSendersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	allowedSourceDomainSenders, pageRes, err := q.keeper.GetAllAllowedSourceDomainSendersPaginated(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllowedSourceDomainSendersResponse{AllowedSourceDomainSenders: allowedSourceDomainSenders, Pagination: pageRes}, nil
}
