package delete

import (
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
)

var DeleteOrgCmd = &cobra.Command{
	Use:   "org",
	Short: "delete the org",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		util.NewWargCtlResponse().Delete("/organization/" + id)
	}}

func init() {
	DeleteOrgCmd.Flags().String("id", "", "id of the organization")
	DeleteOrgCmd.MarkFlagRequired("id")
}
