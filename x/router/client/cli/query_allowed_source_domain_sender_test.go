package cli_test

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"

	"google.golang.org/grpc/codes"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/strangelove-ventures/noble/testutil/network"
	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/strangelove-ventures/noble/x/router/client/cli"
	"github.com/strangelove-ventures/noble/x/router/types"
)

func networkWithAllowedSourceDomainObjects(t *testing.T, n uint32) (*network.Network, []types.AllowedSourceDomainSender) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := uint32(0); i < n; i++ {
		address := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		binary.BigEndian.PutUint32(address[28:], i)
		allowedSourceDomainSender := types.AllowedSourceDomainSender{
			DomainId: i,
			Address:  address,
		}
		state.AllowedSourceDomainSenders = append(state.AllowedSourceDomainSenders, allowedSourceDomainSender)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.AllowedSourceDomainSenders
}

func TestShowAllowedSourceDomainSender(t *testing.T) {
	net, objs := networkWithAllowedSourceDomainObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc         string
		sourceDomain uint32
		address      []byte

		args []string
		err  error
		obj  types.AllowedSourceDomainSender
	}{
		{
			desc:         "found",
			sourceDomain: objs[0].DomainId,
			address:      objs[0].Address,
			args:         common,
			obj:          objs[0],
		},
		{
			desc:         "not found",
			sourceDomain: uint32(14),
			address:      []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			args:         common,
			err:          status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.sourceDomain)),
				hex.EncodeToString(tc.address),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowAllowedSourceDomainSender(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryAllowedSourceDomainSenderResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.AllowedSourceDomainSender)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.AllowedSourceDomainSender),
				)
			}
		})
	}
}
