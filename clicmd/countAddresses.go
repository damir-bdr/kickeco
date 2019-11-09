package clicmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
)

var countAddressesCmd = &cobra.Command{
	Use:   "ca",
	Short: "contact countAddresses",
	RunE: func(cmd *cobra.Command, args []string) error {
		return contractCountAddresses()
	},
}

func init() {
	rootCmd.AddCommand(countAddressesCmd)
}

func contractCountAddresses() error {
	count, err := kickCoinContract.CountAddresses(&bind.CallOpts{})
	if err != nil {
		return err
	}

	fmt.Printf("contact countAddresses=%v\n", count)
	return nil
}
