package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble-router/x/router/keeper"
	"github.com/strangelove-ventures/noble-router/x/router/types"
	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNIBCForward(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoreIBCForwardMetadata {
	items := make([]types.StoreIBCForwardMetadata, n)
	for i := range items {
		items[i].SourceDomainSender = strconv.Itoa(i)
		items[i].Nonce = uint64(i)

		keeper.SetIBCForward(ctx, items[i])
	}
	return items
}

func TestIBCForwardGet(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNIBCForward(routerKeeper, ctx, 10)
	for _, item := range items {
		rst, found := routerKeeper.GetIBCForward(
			ctx,
			item.SourceDomainSender,
			item.Nonce,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestIBCForwardRemove(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNIBCForward(routerKeeper, ctx, 10)
	for _, item := range items {
		routerKeeper.DeleteIBCForward(
			ctx,
			item.SourceDomainSender,
			item.Nonce,
		)
		_, found := routerKeeper.GetIBCForward(
			ctx,
			item.SourceDomainSender,
			item.Nonce,
		)
		require.False(t, found)
	}
}

func TestIBCForwardGetAll(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNIBCForward(routerKeeper, ctx, 10)
	ibcForward := make([]types.StoreIBCForwardMetadata, len(items))
	for i, item := range items {
		ibcForward[i] = item
	}
	require.ElementsMatch(t,
		nullify.Fill(ibcForward),
		nullify.Fill(routerKeeper.GetAllIBCForwards(ctx)),
	)
}
