package clicmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
)

var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "contact owner",
	RunE: func(cmd *cobra.Command, args []string) error {
		return contractOwner()
	},
}

func init() {
	rootCmd.AddCommand(ownerCmd)
}

func contractOwner() error {
	owner, err := kickCoinContract.Owner(&bind.CallOpts{})
	if err != nil {
		return err
	}

	fmt.Printf("contact owner: %s\n", owner.Hex())
	return nil
}
