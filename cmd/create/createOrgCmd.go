package create

import (
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
)

var CrateOrgCmd = &cobra.Command{
	Use:   "org",
	Short: "create the org",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		payload := map[string]interface{}{"name": name, "description": description}
		util.NewWargCtlResponse().Post("/organization/", payload)
	}}

func init() {
	CrateOrgCmd.Flags().String("name", "", "name of the organization")
	CrateOrgCmd.Flags().String("description", "", "description of the organization")
}
