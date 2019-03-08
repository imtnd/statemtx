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
	row = sheet.Row(0)
	cell = row.AddCell()
	cell = row.AddCell()
	cell.Value = "State"
	for i, _ := range sm.States {
		cell = row.AddCell()
		cell.SetValue(i)
	}
	err = file.Save("./sample.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
