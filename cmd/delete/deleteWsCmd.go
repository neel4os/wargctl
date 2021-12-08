package delete

import (
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
)

var DeleteWsCmd = &cobra.Command{
	Use:   "workspace",
	Short: "delete the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		util.NewWargCtlResponse().Delete("/workspace/" + id)
	}}

func init() {
	DeleteWsCmd.Flags().String("id", "", "id of the workspace")
	DeleteWsCmd.MarkFlagRequired("id")
}
