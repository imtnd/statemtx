package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func makeXlsx(sm StateMachine) {
	fmt.Println(sm)

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "test"
	err = file.Save("./sample.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
