package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/noble/x/router/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{
						SourceDomain: 0,
						Channel:      "channel-0",
						Port:         "port-0",
					},
					{
						SourceDomain: 1,
						Channel:      "channel-1",
						Port:         "port-1",
					},
				},
				Mints: []types.Mint{
					{
						SourceDomain:       0,
						SourceDomainSender: []byte("12345678901234567890123456789012"),
						Nonce:              0,
						Amount: &sdk.Coin{
							Denom:  "uusdc",
							Amount: sdk.NewInt(1),
						},
						DestinationDomain: 4,
						MintRecipient:     "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a",
					},
					{
						SourceDomain:       1,
						SourceDomainSender: []byte("12345678901234567890123456789012"),
						Nonce:              1,
						Amount: &sdk.Coin{
							Denom:  "uusdc",
							Amount: sdk.NewInt(2),
						},
						DestinationDomain: 4,
						MintRecipient:     "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a",
					},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{
						SourceDomain: 0,
						Metadata: &types.IBCForwardMetadata{
							Nonce:               0,
							DestinationReceiver: "1234",
							Channel:             "channel-1",
							Port:                "port-1",
						},
					},
					{
						SourceDomain: 1,
						Metadata: &types.IBCForwardMetadata{
							Nonce:               1,
							DestinationReceiver: "1234",
							Channel:             "channel-1",
							Port:                "port-1",
						},
					},
				},
				AllowedSourceDomainSenders: []types.AllowedSourceDomainSender{
					{
						DomainId: 0,
						Address:  []byte("12345678901234567890123456789012"),
					},
					{
						DomainId: 1,
						Address:  []byte("12345678901234567890123456789012"),
					},
				},
			},
			valid: true,
		},
		{
			desc: "duplicated mints",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomain: 0, Port: "transfer", Channel: "channel-1"},
					{SourceDomain: 1, Port: "transfer", Channel: "channel-2"},
				},
				Mints: []types.Mint{
					{Amount: &sdk.Coin{Amount: sdk.OneInt(), Denom: "uusdc"}, MintRecipient: "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a", SourceDomain: 1, SourceDomainSender: []byte("12345678901234567890123456789012"), DestinationDomain: 4},
					{Amount: &sdk.Coin{Amount: sdk.OneInt(), Denom: "uusdc"}, MintRecipient: "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a", SourceDomain: 1, SourceDomainSender: []byte("12345678901234567890123456789012"), DestinationDomain: 4},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomain: 0, Metadata: &types.IBCForwardMetadata{Nonce: 0, Port: "transfer", Channel: "channel-1", DestinationReceiver: "1234"}},
					{SourceDomain: 1, Metadata: &types.IBCForwardMetadata{Nonce: 0, Port: "transfer", Channel: "channel-1", DestinationReceiver: "1234"}},
				},
				AllowedSourceDomainSenders: []types.AllowedSourceDomainSender{
					{DomainId: 0, Address: []byte("12345678901234567890123456789012")},
					{DomainId: 1, Address: []byte("12345678901234567890123456789012")},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated in flight packets",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomain: 1, Port: "transfer", Channel: "channel-1"},
					{SourceDomain: 1, Port: "transfer", Channel: "channel-1"},
				},
				Mints: []types.Mint{
					{Amount: &sdk.Coin{Amount: sdk.OneInt(), Denom: "uusdc"}, MintRecipient: "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a", SourceDomain: 0, SourceDomainSender: []byte("12345678901234567890123456789012"), DestinationDomain: 4},
					{Amount: &sdk.Coin{Amount: sdk.OneInt(), Denom: "uusdc"}, MintRecipient: "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a", SourceDomain: 1, SourceDomainSender: []byte("12345678901234567890123456789012"), DestinationDomain: 4},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomain: 0, Metadata: &types.IBCForwardMetadata{Nonce: 0, Port: "transfer", Channel: "channel-1", DestinationReceiver: "1234"}},
					{SourceDomain: 1, Metadata: &types.IBCForwardMetadata{Nonce: 0, Port: "transfer", Channel: "channel-1", DestinationReceiver: "1234"}},
				},
				AllowedSourceDomainSenders: []types.AllowedSourceDomainSender{
					{DomainId: 0, Address: []byte("12345678901234567890123456789012")},
					{DomainId: 1, Address: []byte("12345678901234567890123456789012")},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated ibc forwards",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomain: 0, Port: "transfer", Channel: "channel-1"},
					{SourceDomain: 1, Port: "transfer", Channel: "channel-2"},
				},
				Mints: []types.Mint{
					{Amount: &sdk.Coin{Amount: sdk.OneInt(), Denom: "uusdc"}, MintRecipient: "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a", SourceDomain: 0, SourceDomainSender: []byte("12345678901234567890123456789012"), DestinationDomain: 4},
					{Amount: &sdk.Coin{Amount: sdk.OneInt(), Denom: "uusdc"}, MintRecipient: "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a", SourceDomain: 1, SourceDomainSender: []byte("12345678901234567890123456789012"), DestinationDomain: 4},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomain: 1, Metadata: &types.IBCForwardMetadata{Nonce: 0, Port: "transfer", Channel: "channel-1", DestinationReceiver: "1234"}},
					{SourceDomain: 1, Metadata: &types.IBCForwardMetadata{Nonce: 0, Port: "transfer", Channel: "channel-1", DestinationReceiver: "1234"}},
				},
				AllowedSourceDomainSenders: []types.AllowedSourceDomainSender{
					{DomainId: 0, Address: []byte("12345678901234567890123456789012")},
					{DomainId: 1, Address: []byte("12345678901234567890123456789012")},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated allowed source domain senders",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomain: 0, Port: "transfer", Channel: "channel-1"},
					{SourceDomain: 1, Port: "transfer", Channel: "channel-2"},
				},
				Mints: []types.Mint{
					{Amount: &sdk.Coin{Amount: sdk.OneInt(), Denom: "uusdc"}, MintRecipient: "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a", SourceDomain: 0, SourceDomainSender: []byte("12345678901234567890123456789012"), DestinationDomain: 4},
					{Amount: &sdk.Coin{Amount: sdk.OneInt(), Denom: "uusdc"}, MintRecipient: "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a", SourceDomain: 1, SourceDomainSender: []byte("12345678901234567890123456789012"), DestinationDomain: 4},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomain: 0, Metadata: &types.IBCForwardMetadata{Nonce: 0, Port: "transfer", Channel: "channel-1", DestinationReceiver: "1234"}},
					{SourceDomain: 1, Metadata: &types.IBCForwardMetadata{Nonce: 0, Port: "transfer", Channel: "channel-1", DestinationReceiver: "1234"}},
				},
				AllowedSourceDomainSenders: []types.AllowedSourceDomainSender{
					{DomainId: 0, Address: []byte("12345678901234567890123456789012")},
					{DomainId: 0, Address: []byte("12345678901234567890123456789012")},
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
