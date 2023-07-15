package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/strangelove-ventures/noble-router/x/router/types"
	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/nullify"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestIBCForwardQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RouterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNIBCForward(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetIBCForwardRequest
		response *types.QueryGetIBCForwardResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetIBCForwardRequest{
				SourceContractAddress: msgs[0].SourceDomainSender,
				Nonce:                 msgs[0].Nonce,
			},
			response: &types.QueryGetIBCForwardResponse{IbcForward: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetIBCForwardRequest{
				SourceContractAddress: msgs[1].SourceDomainSender,
				Nonce:                 msgs[1].Nonce,
			},
			response: &types.QueryGetIBCForwardResponse{IbcForward: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetIBCForwardRequest{
				SourceContractAddress: "nothing",
				Nonce:                 uint64(2),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.IBCForward(wctx, tc.request)
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

func TestIBCForwardQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RouterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNIBCForward(keeper, ctx, 5)
	IBCForward := make([]types.StoreIBCForwardMetadata, len(msgs))
	for i, msg := range msgs {
		IBCForward[i] = msg
	}

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllIBCForwardsRequest {
		return &types.QueryAllIBCForwardsRequest{
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
		for i := 0; i < len(IBCForward); i += step {
			resp, err := keeper.IBCForwards(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.IbcForwards), step)
			require.Subset(t,
				nullify.Fill(IBCForward),
				nullify.Fill(resp.IbcForwards),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(IBCForward); i += step {
			resp, err := keeper.IBCForwards(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.IbcForwards), step)
			require.Subset(t,
				nullify.Fill(IBCForward),
				nullify.Fill(resp.IbcForwards),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.IBCForwards(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(IBCForward), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(IBCForward),
			nullify.Fill(resp.IbcForwards),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.IBCForwards(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
