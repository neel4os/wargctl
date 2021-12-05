package update

import (
	"fmt"
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
)

var UpdateOrgCmd = &cobra.Command{
	Use:   "org",
	Short: "update the org",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		id, _ := cmd.Flags().GetString("id")
		isNameEmpty := name == ""
		isDescriptionEmpty := description == ""
		if isNameEmpty && isDescriptionEmpty {
			fmt.Println("Error : name and description both can not be empty")
		}
		payload := make(map[string]interface{})
		if !isNameEmpty {
			payload["name"] = name
		}
		if !isDescriptionEmpty {
			payload["description"] = description
		}
		util.NewWargCtlResponse().Update("/organization/"+id, payload)
	}}

func init() {
	UpdateOrgCmd.Flags().String("name", "", "name of the organization")
	UpdateOrgCmd.Flags().String("description", "", "description of the organization")
	UpdateOrgCmd.Flags().String("id", "", "id of the organization")
	UpdateOrgCmd.MarkFlagRequired("id")
}
