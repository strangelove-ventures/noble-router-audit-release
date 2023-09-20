package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/sample"
	"github.com/strangelove-ventures/noble/x/router/keeper"
	"github.com/strangelove-ventures/noble/x/router/types"
	"github.com/stretchr/testify/require"
)

/*
* Happy path
* Authority not set
* Invalid authority
* Allowed source domain sender already found
 */

func TestAddAllowedSourceDomainSenderHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.RouterKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgAddAllowedSourceDomainSender{
		From:     owner,
		DomainId: 16,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	_, err := server.AddAllowedSourceDomainSender(sdk.WrapSDKContext(ctx), &message)
	require.Nil(t, err)

	allowed := testkeeper.IsAllowedSourceDomainSender(ctx, message.DomainId, message.Address)
	require.True(t, allowed)
}

func TestAddAllowedSourceDomainSenderAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.RouterKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgAddAllowedSourceDomainSender{
		From:     sample.AccAddress(),
		DomainId: 16,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	require.Panicsf(t, func() {
		_, _ = server.AddAllowedSourceDomainSender(sdk.WrapSDKContext(ctx), &message)
	}, "router owner not found in state")
}

func TestAddAllowedSourceDomainSenderInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.RouterKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgAddAllowedSourceDomainSender{
		From:     "not the authority address",
		DomainId: 16,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	_, err := server.AddAllowedSourceDomainSender(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot add allowed source domain senders")
}

func TestAddAllowedSourceDomainSenderTokenMessengerAlreadyFound(t *testing.T) {
	testkeeper, ctx := keepertest.RouterKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	domainID := uint32(3)
	address := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD}

	testkeeper.AddAllowedSourceDomainSender(ctx, 3, address)

	message := types.MsgAddAllowedSourceDomainSender{
		From:     owner,
		DomainId: domainID,
		Address:  address,
	}

	_, err := server.AddAllowedSourceDomainSender(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrAllowedSourceDomainSenderAlreadyFound, err)
	require.Contains(t, err.Error(), "this source domain sender is already allowed")
}
