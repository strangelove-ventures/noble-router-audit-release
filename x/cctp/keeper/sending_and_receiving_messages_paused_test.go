package keeper_test

import (
	"testing"

	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/x/cctp/types"
	"github.com/stretchr/testify/require"
)

func TestSendingAndReceivingMessagesPaused(t *testing.T) {
	keeper, ctx := keepertest.CctpKeeper(t)

	_, found := keeper.GetSendingAndReceivingMessagesPaused(ctx)
	require.False(t, found)

	paused := types.SendingAndReceivingMessagesPaused{Paused: true}
	keeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	isPaused, found := keeper.GetSendingAndReceivingMessagesPaused(ctx)
	require.True(t, found)
	require.True(t, isPaused.Paused)

	newPaused := types.SendingAndReceivingMessagesPaused{Paused: false}

	keeper.SetSendingAndReceivingMessagesPaused(ctx, newPaused)

	isPaused, found = keeper.GetSendingAndReceivingMessagesPaused(ctx)
	require.True(t, found)
	require.False(t, isPaused.Paused)
}
