package types

import (
	"bytes"
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channelTypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

// MESSAGE FORMAT: <nonce> <sender> <channel> <bech32 prefix> <recipient> <memo>
const (
	NonceIndex   = 0
	NonceLength  = 8
	SenderIndex  = NonceIndex + NonceLength
	SenderLength = 32

	ChannelIndex    = SenderIndex + SenderLength
	ChannelLength   = 8
	PrefixIndex     = ChannelIndex + ChannelLength
	PrefixLength    = 32
	RecipientIndex  = PrefixIndex + PrefixLength
	RecipientLength = 32
	MemoIndex       = RecipientIndex + RecipientLength
)

// Parse parses a byte array into a IBCForwardMetadata struct.
func (m *IBCForwardMetadata) Parse(bz []byte) (*IBCForwardMetadata, error) {
	if len(bz) < MemoIndex {
		return m, ErrDecodingIBCForward
	}

	cutset := string(byte(0))

	m.Nonce = binary.BigEndian.Uint64(bz[NonceIndex:SenderIndex])
	m.Port = "transfer"
	m.Channel = channelTypes.FormatChannelIdentifier(
		binary.BigEndian.Uint64(bz[ChannelIndex:PrefixIndex]),
	)

	prefix := string(bytes.TrimLeft(bz[PrefixIndex:RecipientIndex], cutset))
	recipient := bytes.TrimLeft(bz[RecipientIndex:MemoIndex], cutset)
	m.DestinationReceiver = sdk.MustBech32ifyAddressBytes(prefix, recipient)

	m.Memo = string(bz[MemoIndex:])

	return m, nil
}

// Bytes parses a IBCForwardMetadata struct into a byte array.
func (m *IBCForwardMetadata) Bytes(prefix string) (res []byte, err error) {
	nonceBz := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBz, m.Nonce)

	channelBz := make([]byte, 8)
	rawChannel, err := channelTypes.ParseChannelSequence(m.Channel)
	if err != nil {
		return
	}
	binary.BigEndian.PutUint64(channelBz, rawChannel)

	prefixBz := make([]byte, 32)
	rawPrefix := []byte(prefix)
	copy(prefixBz[32-len(rawPrefix):], rawPrefix)

	recipientBz := make([]byte, 32)
	rawRecipient, err := sdk.GetFromBech32(m.DestinationReceiver, prefix)
	if err != nil {
		return
	}
	copy(recipientBz[32-len(rawRecipient):], rawRecipient)

	res = append(res, nonceBz...)
	res = append(res, make([]byte, 32)...) // sender
	res = append(res, channelBz...)
	res = append(res, prefixBz...)
	res = append(res, recipientBz...)
	res = append(res, []byte(m.Memo)...)
	return
}
