package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/strangelove-ventures/noble/x/cctp/types"
)

func TestRemoteTokenMessengerQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CctpKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRemoteTokenMessengers(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryRemoteTokenMessengerRequest
		response *types.QueryRemoteTokenMessengerResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryRemoteTokenMessengerRequest{
				DomainId: 0,
			},
			response: &types.QueryRemoteTokenMessengerResponse{RemoteTokenMessenger: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryRemoteTokenMessengerRequest{
				DomainId: 1,
			},
			response: &types.QueryRemoteTokenMessengerResponse{RemoteTokenMessenger: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryRemoteTokenMessengerRequest{
				DomainId: 123,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RemoteTokenMessenger(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestRemoteTokenMessengerQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CctpKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRemoteTokenMessengers(keeper, ctx, 5)
	RemoteTokenMessenger := make([]types.RemoteTokenMessenger, len(msgs))
	copy(RemoteTokenMessenger, msgs)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryRemoteTokenMessengersRequest {
		return &types.QueryRemoteTokenMessengersRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(RemoteTokenMessenger); i += step {
			resp, err := keeper.RemoteTokenMessengers(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RemoteTokenMessengers), step)
			require.Subset(t,
				nullify.Fill(RemoteTokenMessenger),
				nullify.Fill(resp.RemoteTokenMessengers),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(RemoteTokenMessenger); i += step {
			resp, err := keeper.RemoteTokenMessengers(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RemoteTokenMessengers), step)
			require.Subset(t,
				nullify.Fill(RemoteTokenMessenger),
				nullify.Fill(resp.RemoteTokenMessengers),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RemoteTokenMessengers(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(RemoteTokenMessenger), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(RemoteTokenMessenger),
			nullify.Fill(resp.RemoteTokenMessengers),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RemoteTokenMessengers(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
