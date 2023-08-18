package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateOwner = "update_owner"

var _ sdk.Msg = &MsgUpdateOwner{}

func NewMsgUpdateOwner(from string, newOwner string) *MsgUpdateOwner {
	return &MsgUpdateOwner{
		From:     from,
		NewOwner: newOwner,
	}
}

func (msg *MsgUpdateOwner) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOwner) Type() string {
	return TypeMsgUpdateOwner
}

func (msg *MsgUpdateOwner) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgUpdateOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	return nil
}
