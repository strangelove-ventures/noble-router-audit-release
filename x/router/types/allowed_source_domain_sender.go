package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

const SourceDomainSenderLen = 32

func (a AllowedSourceDomainSender) Validate() error {
	if len(a.Address) != SourceDomainSenderLen {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "source domain sender address must be %d bytes", SourceDomainSenderLen)
	}
	return nil
}
