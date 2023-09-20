package types

import (
	"testing"

	"github.com/strangelove-ventures/noble/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRemoveAllowedSourceDomainSender_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveAllowedSourceDomainSender
		err  error
	}{
		{
			name: "invalid from",
			msg: MsgRemoveAllowedSourceDomainSender{
				From:     "invalid_address",
				DomainId: uint32(1),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "too short",
			msg: MsgRemoveAllowedSourceDomainSender{
				From:     sample.AccAddress(),
				DomainId: uint32(123),
				Address:  []byte{0x1, 0x23},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "valid",
			msg: MsgRemoveAllowedSourceDomainSender{
				From:     sample.AccAddress(),
				DomainId: uint32(123),
				Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
