package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveAllowedSourceDomainSender{}

func NewMsgRemoveAllowedSourceDomainSender(from string, domainID uint32, address []byte) *MsgRemoveAllowedSourceDomainSender {
	return &MsgRemoveAllowedSourceDomainSender{
		From:     from,
		DomainId: domainID,
		Address:  address,
	}
}

func (msg *MsgRemoveAllowedSourceDomainSender) AllowedSourceDomainSender() AllowedSourceDomainSender {
	return AllowedSourceDomainSender{
		DomainId: msg.DomainId,
		Address:  msg.Address,
	}
}

func (msg *MsgRemoveAllowedSourceDomainSender) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgRemoveAllowedSourceDomainSender) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	if err := msg.AllowedSourceDomainSender().Validate(); err != nil {
		return err
	}
	return nil
}
