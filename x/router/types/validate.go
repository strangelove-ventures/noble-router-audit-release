package types

import (
	"fmt"

	"github.com/circlefin/noble-cctp-router-private/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// Validate ensures that the fields are populated with data that is semantically correct.
func (m *Mint) Validate() error {
	if m.SourceDomainSender == "" {
		return sdkerrors.Wrap(ErrInvalidMint, "the source domain sender cannot be an empty string")
	}

	if m.DestinationDomain != types.NobleDomainId {
		return sdkerrors.Wrapf(ErrInvalidMint, "received an invalid destination domain, expected(%d) got(%d)", types.NobleDomainId, m.DestinationDomain)
	}

	if _, err := sdk.AccAddressFromBech32(m.MintRecipient); err != nil {
		return sdkerrors.Wrapf(ErrInvalidMint, "mint recipient %s is not a valid Noble address", m.MintRecipient)
	}

	if m.Amount == nil {
		return sdkerrors.Wrap(ErrInvalidMint, "amount cannot be nil")
	}

	if err := m.Amount.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidMint, "amount validation error occurred %s", err)
	}

	return nil
}

// Validate ensures that the fields are populated with data that is semantically correct.
func (i *InFlightPacket) Validate() error {
	if i.SourceDomainSender == "" {
		return sdkerrors.Wrap(ErrInvalidInFlightPacket, "the source domain sender cannot be an empty string")
	}

	if err := host.ChannelIdentifierValidator(i.Channel); err != nil {
		return sdkerrors.Wrapf(ErrInvalidInFlightPacket, "invalid channel identifier: %s", err)
	}

	if err := host.PortIdentifierValidator(i.Port); err != nil {
		return sdkerrors.Wrapf(ErrInvalidInFlightPacket, "invalid port identifier: %s", err)
	}

	return nil
}

// Validate ensures that the fields are populated with data that is semantically correct.
func (i *StoreIBCForwardMetadata) Validate() error {
	if i.SourceDomainSender == "" {
		return sdkerrors.Wrap(ErrInvalidStoreForwardMetadata, "the source domain sender cannot be an empty string")
	}

	if err := i.Metadata.Validate(); err != nil {
		return sdkerrors.Wrap(ErrInvalidStoreForwardMetadata, err.Error())
	}

	return nil
}

// Validate ensures that the fields are populated with data that is semantically correct.
func (i *IBCForwardMetadata) Validate() error {
	if i.DestinationReceiver == "" {
		return fmt.Errorf("the destination receiver cannot be an empty string")
	}

	if err := host.ChannelIdentifierValidator(i.Channel); err != nil {
		return fmt.Errorf("invalid channel identifier in IBC forward metadata: %w", err)
	}

	if err := host.PortIdentifierValidator(i.Port); err != nil {
		return fmt.Errorf("invalid port identifier in IBC forward metadata: %w", err)
	}

	return nil
}
