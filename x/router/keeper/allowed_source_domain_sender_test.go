package keeper_test

import (
	"encoding/binary"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/strangelove-ventures/noble/x/router/keeper"
	"github.com/strangelove-ventures/noble/x/router/types"
	"github.com/stretchr/testify/require"
)

func createNAllowedSourceDomainSender(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AllowedSourceDomainSender {
	items := make([]types.AllowedSourceDomainSender, n)
	for i := range items {
		uint32i := uint32(i)
		address := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		binary.BigEndian.PutUint32(address[28:], uint32i)

		items[i].DomainId = uint32i
		items[i].Address = address

		keeper.AddAllowedSourceDomainSender(ctx, uint32i, address)
	}
	return items
}

func TestAllowedSourceDomainSenderGet(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNAllowedSourceDomainSender(routerKeeper, ctx, 10)
	for _, item := range items {
		allowed := routerKeeper.IsAllowedSourceDomainSender(
			ctx,
			item.DomainId,
			item.Address,
		)
		require.True(t, allowed)
	}
}

func TestAllowedSourceDomainSenderRemove(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNAllowedSourceDomainSender(routerKeeper, ctx, 10)
	for _, item := range items {
		routerKeeper.DeleteAllowedSourceDomainSender(
			ctx,
			item.DomainId,
			item.Address,
		)
		allowed := routerKeeper.IsAllowedSourceDomainSender(
			ctx,
			item.DomainId,
			item.Address,
		)
		require.False(t, allowed)
	}
}

func TestAllowedSourceDomainSenderGetAll(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNAllowedSourceDomainSender(routerKeeper, ctx, 10)
	allowedSourceDomainSenders := make([]types.AllowedSourceDomainSender, len(items))
	copy(allowedSourceDomainSenders, items)

	require.ElementsMatch(t,
		nullify.Fill(allowedSourceDomainSenders),
		nullify.Fill(routerKeeper.GetAllowedSourceDomainSenders(ctx)),
	)
}
