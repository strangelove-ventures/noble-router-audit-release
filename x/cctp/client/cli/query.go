package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/strangelove-ventures/noble/x/cctp/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group cctp queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListAttesters())
	cmd.AddCommand(CmdListPerMessageBurnLimits())
	cmd.AddCommand(CmdRemoteTokenMessengers())
	cmd.AddCommand(CmdListTokenPairs())
	cmd.AddCommand(CmdListUsedNonces())
	cmd.AddCommand(CmdShowAttester())
	cmd.AddCommand(CmdRoles())
	cmd.AddCommand(CmdShowBurningAndMintingPaused())
	cmd.AddCommand(CmdShowMaxMessageBodySize())
	cmd.AddCommand(CmdShowNextAvailableNonce())
	cmd.AddCommand(CmdShowPerMessageBurnLimit())
	cmd.AddCommand(CmdShowSendingAndReceivingMessagesPaused())
	cmd.AddCommand(CmdShowSignatureThreshold())
	cmd.AddCommand(CmdRemoteTokenMessenger())
	cmd.AddCommand(CmdShowTokenPair())
	cmd.AddCommand(CmdShowUsedNonce())
	cmd.AddCommand(CmdBurnMessageVersion())
	cmd.AddCommand(CmdLocalMessageVersion())
	cmd.AddCommand(CmdLocalDomain())

	return cmd
}
