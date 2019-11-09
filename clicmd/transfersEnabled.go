package clicmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
)

var transfersEnabledCmd = &cobra.Command{
	Use:   "te",
	Short: "contact transfersEnabled",
	RunE: func(cmd *cobra.Command, args []string) error {
		return contractTransfersEnabled()
	},
}

func init() {
	rootCmd.AddCommand(transfersEnabledCmd)
}

func contractTransfersEnabled() error {
	enabled, err := kickCoinContract.TransfersEnabled(&bind.CallOpts{})
	if err != nil {
		return err
	}

	fmt.Printf("contact TransfersEnabled: %v\n", enabled)
	return nil
}
