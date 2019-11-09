package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

const GATEWAY = "https://mainnet.infura.io/v3/3b3b55afe03647b8a528ba48a71211eb"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "KickCoin contract interaction",
	Long:  "CLI program for calling and viewing public methods of KickCoin token contract",
}

func init() {
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
