package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/strangelove-ventures/noble-router/x/router/types"
)

func CmdListIBCForwards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-ibc-forwards",
		Short: "lists all IBC Forwards",
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
		Use:   "show-ibc-forward [source-contract-address] [nonce]",
		Short: "shows an IBC Forward",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			sourceContractAddress := args[0]
			nonceRaw := args[1]
			nonce, err := strconv.ParseUint(nonceRaw, 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetIBCForwardRequest{
				SourceContractAddress: sourceContractAddress,
				Nonce:                 nonce,
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
