package clicmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
)

var decimalsCmd = &cobra.Command{
	Use:   "decimals",
	Short: "contact decimals",
	RunE: func(cmd *cobra.Command, args []string) error {
		return contractDecimals()
	},
}

func init() {
	rootCmd.AddCommand(decimalsCmd)
}

func contractDecimals() error {
	decimals, err := kickCoinContract.Decimals(&bind.CallOpts{})
	if err != nil {
		return err
	}

	fmt.Printf("contact decimals=%d\n", decimals)
	return nil
}
