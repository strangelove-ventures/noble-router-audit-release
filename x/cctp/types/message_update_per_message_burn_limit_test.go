package types

import (
	"testing"

	"github.com/strangelove-ventures/noble/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdatePerMessageBurnLimit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdatePerMessageBurnLimit
		err  error
	}{
		{
			name: "invalid from",
			msg: MsgUpdatePerMessageBurnLimit{
				From: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid from",
			msg: MsgUpdatePerMessageBurnLimit{
				From: sample.AccAddress(),
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
