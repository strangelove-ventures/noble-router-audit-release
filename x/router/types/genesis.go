package types

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		InFlightPackets: []InFlightPacket{},
		Mints:           []Mint{},
		IbcForwards:     []StoreIBCForwardMetadata{},
		Params:          DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	// Check for duplicated index in InFlightPackets
	inFlightPacketsIndexMap := make(map[string]struct{})
	for _, elem := range gs.InFlightPackets {
		index := hex.EncodeToString(LookupKey(elem.SourceDomain, elem.Nonce))
		if _, ok := inFlightPacketsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for InFlightPackets")
		}
		inFlightPacketsIndexMap[index] = struct{}{}

		// Validate the element to ensure semantic correctness
		if err := elem.Validate(); err != nil {
			return err
		}
	}

	allowedSourceDomainSendersIndexMap := make(map[string]struct{})
	for _, elem := range gs.AllowedSourceDomainSenders {
		index := hex.EncodeToString(SourceDomainSenderKey(elem.DomainId, elem.Address))
		if _, ok := allowedSourceDomainSendersIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for AllowedSourceDomainSenders")
		}
		allowedSourceDomainSendersIndexMap[index] = struct{}{}

		// Validate the element to ensure semantic correctness
		if err := elem.Validate(); err != nil {
			return err
		}
	}

	// Check for duplicated index in mints
	mintsIndexMap := make(map[string]struct{})
	for _, elem := range gs.Mints {
		index := hex.EncodeToString(LookupKey(elem.SourceDomain, elem.Nonce))
		if _, ok := mintsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for Mints")
		}
		mintsIndexMap[index] = struct{}{}

		// Validate the element to ensure semantic correctness
		if err := elem.Validate(); err != nil {
			return err
		}

		// ensure that this mint source domain sender is in AllowedSourceDomainSenders
		allowedSourceDomainSenderKey := hex.EncodeToString(SourceDomainSenderKey(elem.SourceDomain, elem.SourceDomainSender))
		if _, ok := allowedSourceDomainSendersIndexMap[allowedSourceDomainSenderKey]; !ok {
			return fmt.Errorf("mint source domain sender not found in AllowedSourceDomainSenders")
		}
	}

	// Check for duplicated index in ibcForwards
	ibcForwardsIndexMap := make(map[string]struct{})
	for _, elem := range gs.IbcForwards {
		index := hex.EncodeToString(LookupKey(elem.SourceDomain, elem.Metadata.Nonce))
		if _, ok := ibcForwardsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for IBCForwards")
		}
		ibcForwardsIndexMap[index] = struct{}{}

		// Validate the element to ensure semantic correctness
		if err := elem.Validate(); err != nil {
			return err
		}
	}

	if gs.Owner != "" {
		if _, err := sdk.AccAddressFromBech32(gs.Owner); err != nil {
			return err
		}
	}

	return gs.Params.Validate()
}
