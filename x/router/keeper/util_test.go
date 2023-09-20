package keeper_test

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channelTypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/strangelove-ventures/noble/testutil/sample"
	"github.com/strangelove-ventures/noble/x/router/keeper"
	"github.com/strangelove-ventures/noble/x/router/types"
	"github.com/stretchr/testify/require"
)

func TestDecodeIBCForward(t *testing.T) {
	recipient := sample.AccAddress()

	for _, tc := range []struct {
		desc     string
		msg      []byte
		expected types.IBCForwardMetadata
		err      error
	}{
		{
			desc: "Happy path",
			msg:  createMockMetadata(42, "channel-0", sdk.Bech32PrefixAccAddr, recipient, "Hello, World!"),
			expected: types.IBCForwardMetadata{
				Nonce:                42,
				Port:                 "transfer",
				Channel:              "channel-0",
				DestinationReceiver:  recipient,
				Memo:                 "Hello, World!",
				TimeoutInNanoseconds: 0,
			},
		},
		{
			desc: "invalid",
			msg:  []byte("not a valid ibc forward message"),
			err:  types.ErrDecodingIBCForward,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := new(types.IBCForwardMetadata).Parse(tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.expected),
					nullify.Fill(result),
				)
			}
		})
	}
}

func bytesFromMessage(msg keeper.Message) []byte {
	result := make([]byte, keeper.MessageBodyIndex+len(msg.MessageBody))

	binary.BigEndian.PutUint32(result[keeper.VersionIndex:keeper.SourceDomainIndex], msg.Version)
	binary.BigEndian.PutUint32(result[keeper.SourceDomainIndex:keeper.DestinationDomainIndex], msg.SourceDomain)
	binary.BigEndian.PutUint32(result[keeper.DestinationDomainIndex:keeper.NonceIndex], msg.DestinationDomain)
	binary.BigEndian.PutUint64(result[keeper.NonceIndex:keeper.SenderIndex], msg.Nonce)

	copyBytes(msg.Sender, &result, 0, keeper.SenderIndex, keeper.Bytes32Len)
	copyBytes(msg.Recipient, &result, 0, keeper.RecipientIndex, keeper.Bytes32Len)
	copyBytes(msg.DestinationCaller, &result, 0, keeper.DestinationCallerIndex, keeper.Bytes32Len)
	copyBytes(msg.MessageBody, &result, 0, keeper.MessageBodyIndex, len(msg.MessageBody))

	return result
}

func bytesFromBurnMessage(msg keeper.BurnMessage) []byte {
	result := make([]byte, keeper.BurnMessageLen)

	binary.BigEndian.PutUint32(result[keeper.VersionIndex:keeper.BurnTokenIndex], msg.Version)
	amountBytes := uint256ToBytes(&msg.Amount)
	copy(result[keeper.AmountIndex:keeper.MsgSenderIndex], amountBytes[:])

	copyBytes(msg.BurnToken, &result, 0, keeper.BurnTokenIndex, keeper.BurnTokenLen)
	copyBytes(msg.MintRecipient, &result, 0, keeper.MintRecipientIndex, keeper.MintRecipientLen)
	copyBytes(msg.MessageSender, &result, 0, keeper.MsgSenderIndex, keeper.MsgSenderLen)

	return result
}

func copyBytes(src []byte, dest *[]byte, srcStartIndex int, destStartIndex int, length int) {
	for i := 0; i < length; i++ {
		(*dest)[destStartIndex+i] = src[srcStartIndex+i]
	}
}

func fillByteArray(start int, n int) []byte {
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = byte(start + i + 1)
	}
	return res
}

// Write uint256 to byte array in big-endian format
func uint256ToBytes(value *big.Int) []byte {
	// Create a buffer
	buf := new(bytes.Buffer)
	// Write the value into the buffer using big-endian byte order
	buf.Write(value.Bytes())

	// Pad the byte slice if it's not 32 bytes (256 bits) long
	padding := make([]byte, 32-len(buf.Bytes()))

	arr := make([]byte, 32)
	copy(arr, padding)
	copy(arr[len(padding):], buf.Bytes())

	return arr
}

func createMockMetadata(nonce uint64, channel string, prefix string, recipient string, memo string) (res []byte) {
	nonceBz := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBz, nonce)

	channelBz := make([]byte, 8)
	rawChannel, _ := channelTypes.ParseChannelSequence(channel)
	binary.BigEndian.PutUint64(channelBz, rawChannel)

	prefixBz := make([]byte, 32)
	rawPrefix := []byte(prefix)
	copy(prefixBz[32-len(rawPrefix):], rawPrefix)

	recipientBz := make([]byte, 32)
	rawRecipient, _ := sdk.GetFromBech32(recipient, prefix)
	copy(recipientBz[32-len(rawRecipient):], rawRecipient)

	res = append(res, nonceBz...)
	res = append(res, make([]byte, 32)...) // sender
	res = append(res, channelBz...)
	res = append(res, prefixBz...)
	res = append(res, recipientBz...)
	res = append(res, []byte(memo)...)
	return
}
