package get

import (
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
)

var GetExpCmd = &cobra.Command{
	Use:   "experiment",
	Short: "get the experiment",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		path := "/experiment/"
		if id != "" {
			path = path + id
		}
		util.NewWargCtlResponse().Get(path)
	}}

func init() {
	GetExpCmd.Flags().String("id", "", "id of the experiment")
}
