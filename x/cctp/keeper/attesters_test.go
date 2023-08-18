package keeper_test

import (
	"strconv"
	"testing"

	"github.com/strangelove-ventures/noble/x/cctp/keeper"

	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/nullify"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble/x/cctp/types"
	"github.com/stretchr/testify/require"
)

func createNAttesters(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Attester {
	items := make([]types.Attester, n)
	for i := range items {
		items[i].Attester = "Attester" + strconv.Itoa(i)
		keeper.SetAttester(ctx, items[i])
	}
	return items
}

func TestAttestersGet(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNAttesters(cctpKeeper, ctx, 10)
	for _, item := range items {
		attester, found := cctpKeeper.GetAttester(ctx,
			item.Attester,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&attester),
		)
	}
}

func TestAttestersRemove(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNAttesters(cctpKeeper, ctx, 10)
	for _, item := range items {
		cctpKeeper.DeleteAttester(ctx, item.Attester)
		_, found := cctpKeeper.GetAttester(ctx, item.Attester)
		require.False(t, found)
	}
}

func TestAttestersGetAll(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNAttesters(cctpKeeper, ctx, 10)
	denom := make([]types.Attester, len(items))
	copy(denom, items)
	require.ElementsMatch(t,
		nullify.Fill(denom),
		nullify.Fill(cctpKeeper.GetAllAttesters(ctx)),
	)
}
