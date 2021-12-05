package get

import (
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
)

var GetOrgCmd = &cobra.Command{
	Use:   "org",
	Short: "get the org",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		path := "/organization/"
		if id != "" {
			path = path + id
		}
		util.NewWargCtlResponse().Get(path)
	}}

func init() {
	GetOrgCmd.Flags().String("id", "", "id of the organization")
}
