package cmd

import (
	"fmt"
	"github.com/neel4os/wargctl/cmd/create"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a resource when resource is defined",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: must also specify a resource. Options are org")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(create.CrateOrgCmd)
}
