package create

import (
	"encoding/json"
	"fmt"
	"github.com/neel4os/wargctl/util"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func read_json_file(filepath string) interface{} {
	var experiment interface{}
	content, err := ioutil.ReadFile(filepath)
	if err == nil {
		jsonerr := json.Unmarshal(content, &experiment)
		if jsonerr == nil {
			return experiment
		}
		panic(jsonerr)
	}
	panic(err)
}

func ErrorHandling(filepath string) {
	r := recover()
	if r != nil {
		fmt.Println("Error loading in " + filepath + " because")
		fmt.Println(r)
	}
}

var CreateExpCmd = &cobra.Command{
	Use:   "experiment",
	Short: "create the experiment",
	Run: func(cmd *cobra.Command, args []string) {
		filepath, _ := cmd.Flags().GetString("filepath")
		workspaceId, _ := cmd.Flags().GetString("workspaceId")
		defer ErrorHandling(filepath)
		content := read_json_file(filepath)
		payload := map[string]interface{}{"experiment": content, "workspaceId": workspaceId}
		util.NewWargCtlResponse().Post("/experiment/", payload)
		fmt.Println(payload)
	}}

func init() {
	CreateExpCmd.Flags().String("filepath", "", "name of the experiment file")
	CreateExpCmd.Flags().String("workspaceId", "", "id of the parent workspace")
	CreateExpCmd.MarkFlagRequired("workspaceId")
	CreateExpCmd.MarkFlagRequired("filepath")
}
