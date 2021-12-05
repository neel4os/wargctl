package cmd

import (
	"fmt"
	"github.com/neel4os/wargctl/cmd/delete"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a resource when resource is defined",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: must also specify a resource. Options are org")
	},
}

func init() {
	rootCmd.AddCommand(DeleteCmd)
	DeleteCmd.AddCommand(delete.DeleteOrgCmd)
}
