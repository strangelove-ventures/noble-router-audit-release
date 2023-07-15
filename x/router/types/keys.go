package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// ModuleName defines the module name
	ModuleName = "router"

	// StoreKey defines the primary module store key
	StoreKey = "router"

	// RouterKey defines the module's message routing key
	RouterKey = StoreKey

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_" + StoreKey

	IBCForwardKeyPrefix     = "IBCForward/value/"
	InFlightPacketKeyPrefix = "InFlightPacket/value/"
	MintKeyPrefix           = "Mint/value/"
)

func LookupKey(sourceDomainSender string, nonce uint64) []byte {

	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, nonce)

	combinedBytes := append(nonceBytes, []byte(sourceDomainSender)...)
	hashedKey := crypto.Keccak256(combinedBytes)

	return append(hashedKey, []byte("/")...)
}

func IBCForwardPrefix(p string) []byte {
	return []byte(p)
}

func InFlightPacketPrefix(p string) []byte {
	return []byte(p)
}

func MintPrefix(p string) []byte {
	return []byte(p)
}

func InFlightPacketKey(channelID, portID string, sequence uint64) []byte {
	return []byte(fmt.Sprintf("%s/%s/%d", channelID, portID, sequence))
}
