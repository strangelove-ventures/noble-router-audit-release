package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/strangelove-ventures/noble/x/router/types"
)

func CmdListIBCForwards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-ibc-forwards",
		Short: "lists all IBC forwards",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllIBCForwardsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.IBCForwards(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowIBCForward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-ibc-forward [source-domain] [nonce]",
		Short: "shows an IBC forward",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			sourceDomain, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}
			nonce, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetIBCForwardRequest{
				SourceDomain: uint32(sourceDomain),
				Nonce:        nonce,
			}

			res, err := queryClient.IBCForward(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
