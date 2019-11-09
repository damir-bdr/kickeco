package clicmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
)

var standartCmd = &cobra.Command{
	Use:   "standart",
	Short: "contact standart",
	RunE: func(cmd *cobra.Command, args []string) error {
		return contractStandart()
	},
}

func init() {
	rootCmd.AddCommand(standartCmd)
}

func contractStandart() error {
	standart, err := kickCoinContract.Standard(&bind.CallOpts{})
	if err != nil {
		return err
	}

	fmt.Printf("contact standart: %s\n", standart)
	return nil
}
