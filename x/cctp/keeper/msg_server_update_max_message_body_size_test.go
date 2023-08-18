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
 */

func TestUpdateMaxMessageBodySizeHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgUpdateMaxMessageBodySize{
		From:        owner,
		MessageSize: uint64(1023),
	}
	_, err := server.UpdateMaxMessageBodySize(sdk.WrapSDKContext(ctx), &message)
	require.Nil(t, err)

	actual, found := testkeeper.GetMaxMessageBodySize(ctx)
	require.True(t, found)
	require.Equal(t, message.MessageSize, actual.Amount)
}

func TestUpdateMaxMessageBodySizeAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgUpdateMaxMessageBodySize{
		From:        sample.AccAddress(),
		MessageSize: uint64(1023),
	}
	_, err := server.UpdateMaxMessageBodySize(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrAuthorityNotSet, err)
	require.Contains(t, err.Error(), "authority not set")
}

func TestUpdateMaxMessageBodySizeInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgUpdateMaxMessageBodySize{
		From:        "not the authority",
		MessageSize: uint64(1023),
	}
	_, err := server.UpdateMaxMessageBodySize(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot update the max message body size")
}
