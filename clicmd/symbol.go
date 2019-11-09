package clicmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
)

var symbolCmd = &cobra.Command{
	Use:   "symbol",
	Short: "contact symbol",
	RunE: func(cmd *cobra.Command, args []string) error {
		return contractSymbol()
	},
}

func init() {
	rootCmd.AddCommand(symbolCmd)
}

func contractSymbol() error {
	symbol, err := kickCoinContract.Symbol(&bind.CallOpts{})
	if err != nil {
		return err
	}

	fmt.Printf("contact symbol: %s\n", symbol)
	return nil
}
