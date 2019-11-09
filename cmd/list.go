package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of the tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		filterDone, err := cmd.Flags().GetBool("todo_only")
		if err != nil {
			return err
		}
		return list(context.Background(), filterDone)
	},
}

func init() {
	listCmd.Flags().BoolP("todo_only", "t", false, "Show only tasks that are not completed yet")
	rootCmd.AddCommand(listCmd)
}

func list(ctx context.Context, filterDone bool) error {
	fmt.Println("list", filterDone)
	return nil
}
