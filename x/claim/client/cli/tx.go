package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmosquad-labs/squad/x/claim/types"
)

// GetTxCmd returns the transaction commands for the module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transConditions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewClaimCmd(),
	)

	return cmd
}

func NewClaimCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim [airdrop-id] [action-type]",
		Args:  cobra.ExactArgs(2),
		Short: "Claim the claimable amount with an action type",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Claim the claimable amount with an action type. 
There are 3 different types of action type. Reference the examples below.

Example:
$ %s tx %s claim 1 deposit --from mykey
$ %s tx %s claim 1 swap --from mykey
$ %s tx %s claim 1 farming --from mykey
`,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			airdropId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			actionTyp := NormalizeConditionType(args[1])
			if actionTyp == types.ConditionTypeUnspecified {
				return fmt.Errorf("unknown action type %s", args[0])
			}

			msg := types.NewMsgClaim(
				airdropId,
				clientCtx.GetFromAddress(),
				actionTyp,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
