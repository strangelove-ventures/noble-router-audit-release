package keeper_test

import (
	"testing"

	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/strangelove-ventures/noble/x/cctp/keeper"

	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble/x/cctp/types"
	"github.com/stretchr/testify/require"
)

func createNUsedNonces(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Nonce {
	items := make([]types.Nonce, n)
	for i := range items {
		items[i].SourceDomain = uint32(i)
		items[i].Nonce = uint64(i)

		keeper.SetUsedNonce(ctx, items[i])
	}
	return items
}

func TestUsedNonceGet(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNUsedNonces(cctpKeeper, ctx, 10)
	for _, item := range items {
		found := cctpKeeper.GetUsedNonce(ctx, item)
		require.True(t, found)
	}
}

func TestUsedNonceGetNotFound(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)

	found := cctpKeeper.GetUsedNonce(ctx,
		types.Nonce{
			SourceDomain: 123,
			Nonce:        0,
		})
	require.False(t, found)
}

func TestUsedNoncesGetAll(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNUsedNonces(cctpKeeper, ctx, 10)
	nonce := make([]types.Nonce, len(items))
	copy(nonce, items)

	require.ElementsMatch(t,
		nullify.Fill(nonce),
		nullify.Fill(cctpKeeper.GetAllUsedNonces(ctx)),
	)
}
