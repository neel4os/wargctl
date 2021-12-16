package util

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
)

type TablePrinterInterface interface {
	Print()
}

type payload []byte
type jsonmap = map[string]interface{}

type TablePrinter struct {
	Payload payload
}

func (receiver TablePrinter) createStructureTable(jnm *jsonmap, fv bool) ([][]string, []string) {
	var header []string
	c := make([][]string, 0)
	keys := getKeys(*jnm)
	for _, v := range keys {
		switch (*jnm)[v].(type) {
		case string:
			d := make([]string, 0)
			if fv {
				d = append(d, v, (*jnm)[v].(string))
			} else {
				d = append(d, (*jnm)[v].(string))
			}
			c = append(c, d)
		case []interface{}:
			sjnm := (*jnm)[v]
			if len(sjnm.([]interface{})) > 0 {
				header = getKeys(sjnm.([]interface{})[0].(jsonmap))
				for _, iter := range sjnm.([]interface{}) {
					//for _, sk := range header {
					d := make([]string, 0)
					t := iter.(jsonmap)
					c1, _ := receiver.createStructureTable(&t, false)
					for _, value := range c1 {
						d = append(d, value[0])
					}
					c = append(c, d)
				}
			} else {
				panic(keys)
			}
		}
	}
	if fv {
		if len(header) == 0 {
			header = []string{"Field", "Value"}
		}
	}
	return c, header
}

func (receiver TablePrinter) Print() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("No", r.([]string)[0], "exist.")
		}
	}()
	jsonPayload := receiver.convertMap()
	arr, headers := receiver.createStructureTable(jsonPayload, true)
	table := receiver.setTable(headers)
	table.AppendBulk(arr)
	table.Render()
}

func (receiver TablePrinter) convertMap() *jsonmap {
	var jsonPayload jsonmap
	json.Unmarshal(receiver.Payload, &jsonPayload)
	return &jsonPayload
}

func (receiver TablePrinter) setTable(header []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	return table
}

func getKeys(m jsonmap) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
