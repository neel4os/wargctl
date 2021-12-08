package create

import (
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
)

var CreateWsCmd = &cobra.Command{
	Use:   "workspace",
	Short: "create the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		organizationId, _ := cmd.Flags().GetString("organizationId")
		payload := map[string]interface{}{"name": name, "description": description, "organizationId": organizationId}
		util.NewWargCtlResponse().Post("/workspace/", payload)
	}}

func init() {
	CreateWsCmd.Flags().String("name", "", "name of the organization")
	CreateWsCmd.Flags().String("description", "", "description of the organization")
	CreateWsCmd.Flags().String("organizationId", "", "id of the parent organization")
	CreateWsCmd.MarkFlagRequired("organizationId")
}
