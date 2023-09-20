package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/strangelove-ventures/noble/x/router/types"
)

func CmdListAllowedSourceDomainSenders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-allowed-source-domain-senders",
		Short: "lists all allowed source domain senders",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllowedSourceDomainSendersRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllowedSourceDomainSenders(context.Background(), params)
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

func CmdShowAllowedSourceDomainSender() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-allowed-source-domain-sender [source-domain] [address]",
		Short: "shows an allowed source domain sender",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			sourceDomain, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}
			address := common.FromHex(args[1])

			params := &types.QueryAllowedSourceDomainSenderRequest{
				DomainId: uint32(sourceDomain),
				Address:  address,
			}

			res, err := queryClient.AllowedSourceDomainSender(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
