package keeper_test

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/strangelove-ventures/noble/x/router/keeper"
	"github.com/strangelove-ventures/noble/x/router/types"
	"strconv"
	"testing"

	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestInvalidOuterMessage(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	msg := []byte("not a valid message")
	err := routerKeeper.HandleMessage(ctx, msg)

	require.ErrorIs(t, err, sdkerrors.Wrap(types.ErrDecodingMessage, "error decoding message"))
}

func TestInvalidOuterMessageBody(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	msg := bytesFromMessage(keeper.Message{
		Version:           1,
		SourceDomain:      2,
		DestinationDomain: 3,
		Nonce:             4,
		Sender:            fillByteArray(0, 32),
		Recipient:         fillByteArray(32, 32),
		DestinationCaller: fillByteArray(64, 32),
		MessageBody:       []byte(""),
	})

	err := routerKeeper.HandleMessage(ctx, msg)
	if err != nil {
		return
	}
}
