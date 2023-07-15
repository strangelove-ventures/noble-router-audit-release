package types_test

import (
	"testing"

	"github.com/strangelove-ventures/noble-router/x/router/types"
	"github.com/strangelove-ventures/noble/testutil/sample"

	"github.com/stretchr/testify/require"
)

var testAddress = sample.AccAddress()

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    false,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomainSender: "0"},
					{SourceDomainSender: "1"},
				},
				Mints: []types.Mint{
					{SourceDomainSender: "0"},
					{SourceDomainSender: "1"},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomainSender: "0"},
					{SourceDomainSender: "1"},
				},
			},
			valid: true,
		},
		{
			desc: "mints is nil",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomainSender: "0"},
					{SourceDomainSender: "1"},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomainSender: "0"},
					{SourceDomainSender: "1"},
				},
			},
			valid: false,
		},
		{
			desc: "mints is nil",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomainSender: "0"},
					{SourceDomainSender: "1"},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomainSender: "0"},
					{SourceDomainSender: "1"},
				},
			},
			valid: false,
		},
		{
			desc: "ibcforwardmetadata is nil",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				Mints: []types.Mint{
					{SourceDomainSender: "0"},
					{SourceDomainSender: "1"},
				},
				InFlightPackets: []types.InFlightPacket{
					{SourceDomainSender: "0"},
					{SourceDomainSender: "1"},
				},
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
