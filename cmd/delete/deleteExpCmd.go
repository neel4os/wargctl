package delete

import (
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
)

var DeleteExpCmd = &cobra.Command{
	Use:   "experiment",
	Short: "delete the experiment",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		util.NewWargCtlResponse().Delete("/experiment/" + id)
	}}

func init() {
	DeleteExpCmd.Flags().String("id", "", "id of the experiment")
	DeleteExpCmd.MarkFlagRequired("id")
}
