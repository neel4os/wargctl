package cmd

import (
	"fmt"
	"github.com/neel4os/wargctl/cmd/get"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a resource when resource is defined",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: must also specify a resource. Options are org")
	},
}

func init() {
	rootCmd.AddCommand(GetCmd)
	GetCmd.AddCommand(get.GetOrgCmd)
}
