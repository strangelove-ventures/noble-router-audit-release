package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/strangelove-ventures/noble/x/cctp/types"
)

func CmdReplaceDepositForBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "replace-deposit-for-burn [original-message] [original-attestation] [new-destination-caller] [new-mint-recipient]",
		Short: "Broadcast message replace-deposit-for-burn",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgReplaceDepositForBurn(
				clientCtx.GetFromAddress().String(),
				[]byte(args[0]),
				[]byte(args[1]),
				[]byte(args[2]),
				[]byte(args[3]),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
