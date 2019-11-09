package clicmd

import (
	"log"

	"github.com/spf13/cobra"
)

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
