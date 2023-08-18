package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/x/cctp/types"
)

func TestBurningAndMintingPaused(t *testing.T) {
	keeper, ctx := keepertest.CctpKeeper(t)

	paused := types.BurningAndMintingPaused{Paused: true}
	keeper.SetBurningAndMintingPaused(ctx, paused)

	isPaused, found := keeper.GetBurningAndMintingPaused(ctx)
	require.True(t, found)
	require.True(t, isPaused.Paused)

	newPaused := types.BurningAndMintingPaused{Paused: false}

	keeper.SetBurningAndMintingPaused(ctx, newPaused)

	isPaused, found = keeper.GetBurningAndMintingPaused(ctx)
	require.True(t, found)
	require.False(t, isPaused.Paused)
}
