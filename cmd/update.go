package cmd

import (
	"fmt"
	"github.com/neel4os/wargctl/cmd/update"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a resource when resource is defined",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: must also specify a resource. Options are org")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(update.UpdateOrgCmd)
}
