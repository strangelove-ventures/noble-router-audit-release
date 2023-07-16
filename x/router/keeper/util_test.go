package keeper_test

import (
	"encoding/binary"
	"github.com/strangelove-ventures/noble/x/router/keeper"
	"github.com/strangelove-ventures/noble/x/router/types"
	"testing"

	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/stretchr/testify/require"
)

func TestDecodeMessage(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		msg      []byte
		expected keeper.Message
		err      error
	}{
		{
			desc: "Happy path",
			msg: bytesFromMessage(keeper.Message{
				Version:           1,
				SourceDomain:      2,
				DestinationDomain: 3,
				Nonce:             4,
				Sender:            fillByteArray(0, 32),
				Recipient:         fillByteArray(32, 32),
				DestinationCaller: fillByteArray(64, 32),
				MessageBody:       []byte("your average run of the mill message body"),
			}),
			expected: keeper.Message{
				Version:           1,
				SourceDomain:      2,
				DestinationDomain: 3,
				Nonce:             4,
				Sender:            fillByteArray(0, 32),
				Recipient:         fillByteArray(32, 32),
				DestinationCaller: fillByteArray(64, 32),
				MessageBody:       []byte("your average run of the mill message body"),
			},
		},
		{
			desc: "invalid",
			msg:  []byte("-1"),
			err:  types.ErrDecodingMessage,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := keeper.DecodeMessage(tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.expected),
					nullify.Fill(*result),
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

	copyBytes(msg.Sender, &result, 0, keeper.SenderIndex, len(msg.Sender))
	copyBytes(msg.Recipient, &result, 0, keeper.RecipientIndex, len(msg.Recipient))
	copyBytes(msg.DestinationCaller, &result, 0, keeper.DestinationCallerIndex, len(msg.DestinationCaller))
	copyBytes(msg.MessageBody, &result, 0, keeper.MessageBodyIndex, len(msg.MessageBody))

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
