package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/strangelove-ventures/noble/x/cctp/types"
)

func CmdShowNextAvailableNonce() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-next-available-nonce",
		Short: "shows the next available nonce",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetNextAvailableNonceRequest{}

			res, err := queryClient.NextAvailableNonce(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
