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

func TestTokenPairQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CctpKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTokenPairs(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTokenPairRequest
		response *types.QueryGetTokenPairResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTokenPairRequest{
				RemoteDomain: 0,
				RemoteToken:  "00",
			},
			response: &types.QueryGetTokenPairResponse{Pair: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTokenPairRequest{
				RemoteDomain: 1,
				RemoteToken:  "01",
			},
			response: &types.QueryGetTokenPairResponse{Pair: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTokenPairRequest{
				RemoteDomain: 123,
				RemoteToken:  "123",
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TokenPair(wctx, tc.request)
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

func TestTokenPairQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CctpKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTokenPairs(keeper, ctx, 5)
	TokenPair := make([]types.TokenPair, len(msgs))
	copy(TokenPair, msgs)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTokenPairsRequest {
		return &types.QueryAllTokenPairsRequest{
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
		for i := 0; i < len(TokenPair); i += step {
			resp, err := keeper.TokenPairs(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TokenPairs), step)
			require.Subset(t,
				nullify.Fill(TokenPair),
				nullify.Fill(resp.TokenPairs),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(TokenPair); i += step {
			resp, err := keeper.TokenPairs(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TokenPairs), step)
			require.Subset(t,
				nullify.Fill(TokenPair),
				nullify.Fill(resp.TokenPairs),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TokenPairs(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(TokenPair), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(TokenPair),
			nullify.Fill(resp.TokenPairs),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TokenPairs(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
