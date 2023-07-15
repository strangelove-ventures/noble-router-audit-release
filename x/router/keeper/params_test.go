package keeper_test

import (
	"testing"

	"github.com/strangelove-ventures/noble-router/x/router/types"
	testkeeper "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RouterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
