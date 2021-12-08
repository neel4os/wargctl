package get

import (
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
)

var GetWsCmd = &cobra.Command{
	Use:   "workspace",
	Short: "get the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		path := "/workspace/"
		if id != "" {
			path = path + id
		}
		util.NewWargCtlResponse().Get(path)
	}}

func init() {
	GetWsCmd.Flags().String("id", "", "id of the workspace")
}
