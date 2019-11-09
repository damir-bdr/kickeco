package clicmd

import (
	"log"

	"github.com/damir-bdr/kickeco/contract"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "KickCoin contract interaction",
	Long:  "CLI program for calling and viewing public methods of KickCoin token contract",
}

var kickCoinContract *contract.CSToken

func Execute(contract *contract.CSToken) {
	kickCoinContract = contract

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
