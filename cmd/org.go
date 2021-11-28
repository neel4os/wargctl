package cmd

import (
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
)

type org struct {
	name        string
	description string
}

// orgCmd represents the org command
var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "Create an Organization",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		payload := map[string]interface{}{"name": name, "description": description}
		util.NewWargCtlResponse().Post("/organization/", payload)
	},
}

func init() {
	createCmd.AddCommand(orgCmd)

	orgCmd.PersistentFlags().String("name", "", "name of the organization")
	orgCmd.PersistentFlags().String("description", "", "description of the organization")

}
