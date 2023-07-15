package keeper

import (
	"encoding/binary"
	"github.com/gogo/protobuf/proto"
	"github.com/strangelove-ventures/noble-router/x/router/types"
)

// TODO copy pasted from github.com/strangelove-ventures/noble-cctp, change to reference that

type BurnMessage struct {
	Version       uint32
	BurnToken     []byte
	MintRecipient []byte
	Amount        uint64
	MessageSender []byte
}

type Message struct {
	Version           uint32
	SourceDomain      uint32
	DestinationDomain uint32
	Nonce             uint64
	Sender            []byte
	Recipient         []byte
	DestinationCaller []byte
	MessageBody       []byte
}

const (
	// Indices of each field in message
	VersionIndex           = 0
	SourceDomainIndex      = 4
	DestinationDomainIndex = 8
	NonceIndex             = 12
	SenderIndex            = 20
	RecipientIndex         = 52
	DestinationCallerIndex = 84
	MessageBodyIndex       = 116

	// Indices of each field in BurnMessage
	BurnMsgVersionIndex = 0
	VersionLen          = 4
	BurnTokenIndex      = 4
	BurnTokenLen        = 32
	MintRecipientIndex  = 36
	MintRecipientLen    = 32
	AmountIndex         = 68
	AmountLen           = 32
	MsgSenderIndex      = 100
	MsgSenderLen        = 32
	// 4 byte version + 32 bytes burnToken + 32 bytes mintRecipient + 32 bytes amount + 32 bytes messageSender
	BurnMessageLen = 132

	NobleMessageVersion = 0
	MessageBodyVersion  = 0
	NobleDomainId       = 4
	Bytes32Len          = 32
)

// TODO test these, make sure indices are correct, errs correctly

func decodeBurnMessage(msg []byte) (BurnMessage, error) {
	message := BurnMessage{
		Version:       binary.BigEndian.Uint32(msg[BurnMsgVersionIndex:BurnTokenIndex]),
		BurnToken:     msg[BurnTokenIndex:MintRecipientIndex],
		MintRecipient: msg[MintRecipientIndex:AmountIndex],
		Amount:        binary.BigEndian.Uint64(msg[AmountIndex:MsgSenderIndex]),
		MessageSender: msg[MsgSenderIndex:BurnMessageLen],
	}

	return message, nil
}

func decodeMessage(msg []byte) (Message, error) {
	message := Message{
		Version:           binary.BigEndian.Uint32(msg[VersionIndex:SourceDomainIndex]),
		SourceDomain:      binary.BigEndian.Uint32(msg[SourceDomainIndex:DestinationDomainIndex]),
		DestinationDomain: binary.BigEndian.Uint32(msg[DestinationDomainIndex:NonceIndex]),
		Nonce:             binary.BigEndian.Uint64(msg[NonceIndex:SenderIndex]),
		Sender:            msg[SenderIndex:RecipientIndex],
		Recipient:         msg[RecipientIndex:DestinationCallerIndex],
		DestinationCaller: msg[DestinationCallerIndex:MessageBodyIndex],
		MessageBody:       msg[MessageBodyIndex:],
	}

	return message, nil
}

func decodeIBCForward(msg []byte) (types.IBCForwardMetadata, error) {
	var res types.IBCForwardMetadata
	if err := proto.Unmarshal(msg, &res); err != nil {
		return types.IBCForwardMetadata{}, err
	}

	return res, nil
}
