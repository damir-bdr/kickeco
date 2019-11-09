package clicmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
)

var totalSupplyCmd = &cobra.Command{
	Use:   "totalsup",
	Short: "contact get totalSupply",
	RunE: func(cmd *cobra.Command, args []string) error {
		return contractTotalsupply()
	},
}

func init() {
	rootCmd.AddCommand(totalSupplyCmd)
}

func contractTotalsupply() error {
	totalSupply, err := kickCoinContract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return err
	}

	fmt.Printf("contact TotalSupply: %v\n", totalSupply)
	return nil
}
