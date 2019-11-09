package clicmd

import (
	"fmt"

	"github.com/damir-bdr/kickeco/contract"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List KickCoin token funcs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listCSTokenFuncs()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listCSTokenFuncs() error {
	for _, f := range contract.CSTokenFuncSigs {
		fmt.Println(f)
	}
	return nil
}
