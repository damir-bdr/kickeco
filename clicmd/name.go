package clicmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
)

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "contact name",
	RunE: func(cmd *cobra.Command, args []string) error {
		return contractName()
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)
}

func contractName() error {
	name, err := kickCoinContract.Name(&bind.CallOpts{})
	if err != nil {
		return err
	}

	fmt.Printf("contact name: %s\n", name)
	return nil
}
