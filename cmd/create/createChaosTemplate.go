package create

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

type Experiment struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Method      []string `json:"method"`
}

var CrateTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "create an bare minimum chaos experiment ",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		filename, _ := cmd.Flags().GetString("filename")
		template := Experiment{title, description, make([]string, 0)}
		data, error := json.Marshal(template)
		if error != nil {
			fmt.Println(error)
		}
		err := ioutil.WriteFile(filename, data, 0644)
		if err != nil {
			fmt.Println(filename + " can not be created because")
			fmt.Println(err)
		}
		fmt.Println(filename + "created successfully")
	}}

func init() {
	CrateTemplateCmd.Flags().String("title", "", "title of the experiment")
	CrateTemplateCmd.Flags().String("description", "", "description of the experiment")
	CrateTemplateCmd.Flags().String("filename", "", "name of the experiment file with extention")
	CrateTemplateCmd.MarkFlagRequired("title")
	CrateTemplateCmd.MarkFlagRequired("description")
	CrateTemplateCmd.MarkFlagRequired("filename")
}
