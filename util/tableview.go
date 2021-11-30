package util

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"os"
)

type TablePrinterInterface interface {
	Print()
}

type payload []byte
type jsonmap map[string]string

type TablePrinter struct {
	Payload payload
}

func (receiver TablePrinter) Print() {
	jsonPayload := receiver.convertMap()
	table := receiver.setTable()
	for k, v := range *jsonPayload {
		c := make([]string, 0)
		c = append(c, k, v)
		table.Append(c)
	}
	table.Render()
}

func (receiver TablePrinter) convertMap() *jsonmap {
	var jsonPayload jsonmap
	json.Unmarshal(receiver.Payload, &jsonPayload)
	return &jsonPayload
}

func (receiver TablePrinter) setTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})
	return table
}
