package clicmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var balanceOfCmd = &cobra.Command{
	Use:   "bof",
	Short: "contact BalanceOf for ownerAddr",
	RunE: func(cmd *cobra.Command, args []string) error {
		ownerAddr, err := cmd.Flags().GetString("owner")
		if err != nil {
			return err
		}

		return contractBalanceOfCmd(ownerAddr)
	},
}

func init() {
	balanceOfCmd.Flags().String("owner", "0x387fc6939b5e54b2f11793df05388f9d11942948", "owner address")
	rootCmd.AddCommand(balanceOfCmd)
}

func contractBalanceOfCmd(ownerAddr string) error {
	balance, err := kickCoinContract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(ownerAddr))
	if err != nil {
		return err
	}

	fmt.Printf("contact balance=%v for owner addr: %s\n", balance, ownerAddr)
	return nil
}
