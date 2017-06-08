package main

import ("github.com/crackcomm/go-clitable")

func PrettyPrintMany(headers []string, rows []map[string]interface{}) {
	table := clitable.New(headers)
	for _, row := range rows {
		table.AddRow(row)
	}
	table.Print()
}

func PrettyPrintSingle(row map[string]interface{}) {
	clitable.PrintHorizontal(row)
}