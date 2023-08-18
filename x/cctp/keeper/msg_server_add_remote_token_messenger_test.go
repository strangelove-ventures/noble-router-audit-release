package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/sample"
	"github.com/strangelove-ventures/noble/x/cctp/keeper"
	"github.com/strangelove-ventures/noble/x/cctp/types"
	"github.com/stretchr/testify/require"
)

/*
* Happy path
* Authority not set
* Invalid authority
* Remote token messenger already found
 */

func TestAddRemoteTokenMessengerHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgAddRemoteTokenMessenger{
		From:     owner,
		DomainId: 16,
		Address:  "remote_token_messenger_address",
	}

	_, err := server.AddRemoteTokenMessenger(sdk.WrapSDKContext(ctx), &message)
	require.Nil(t, err)

	actual, found := testkeeper.GetRemoteTokenMessenger(ctx, message.DomainId)
	require.True(t, found)

	require.Equal(t, message.DomainId, actual.DomainId)
	require.Equal(t, message.Address, actual.Address)
}

func TestAddRemoteTokenMessengerAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgAddRemoteTokenMessenger{
		From:     sample.AccAddress(),
		DomainId: 16,
		Address:  "remote_token_messenger_address",
	}

	_, err := server.AddRemoteTokenMessenger(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrAuthorityNotSet, err)
	require.Contains(t, err.Error(), "authority not set")
}

func TestAddRemoteTokenMessengerInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgAddRemoteTokenMessenger{
		From:     "not the authority address",
		DomainId: 16,
		Address:  "remote_token_messenger_address",
	}

	_, err := server.AddRemoteTokenMessenger(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot add remote token messengers")
}

func TestAddRemoteTokenMessengerTokenMessengerAlreadyFound(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	existingRemoteTokenMessenger := types.RemoteTokenMessenger{
		DomainId: 3,
		Address:  sample.AccAddress(),
	}
	testkeeper.SetRemoteTokenMessenger(ctx, existingRemoteTokenMessenger)

	message := types.MsgAddRemoteTokenMessenger{
		From:     owner,
		DomainId: existingRemoteTokenMessenger.DomainId,
		Address:  "remote_token_messenger_address",
	}

	_, err := server.AddRemoteTokenMessenger(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrRemoteTokenMessengerAlreadyFound, err)
	require.Contains(t, err.Error(), "a remote token messenger for this domain already exists")
}
