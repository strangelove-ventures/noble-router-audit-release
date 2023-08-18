package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/strangelove-ventures/noble/x/cctp/types"
)

func TestSignatureThreshold(t *testing.T) {
	keeper, ctx := keepertest.CctpKeeper(t)
	_, found := keeper.GetSignatureThreshold(ctx)
	require.False(t, found)

	SignatureThreshold := types.SignatureThreshold{Amount: 2}
	keeper.SetSignatureThreshold(ctx, SignatureThreshold)

	threshold, found := keeper.GetSignatureThreshold(ctx)
	require.True(t, found)
	require.Equal(t,
		SignatureThreshold,
		nullify.Fill(&threshold),
	)

	newSignatureThreshold := types.SignatureThreshold{Amount: 3}

	keeper.SetSignatureThreshold(ctx, newSignatureThreshold)

	threshold, found = keeper.GetSignatureThreshold(ctx)
	require.True(t, found)
	require.Equal(t,
		newSignatureThreshold,
		nullify.Fill(&threshold),
	)
}
