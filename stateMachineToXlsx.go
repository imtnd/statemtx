package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

func makeXlsx(sm StateMachine) {
	fmt.Println(sm)

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	addStateNameLine(sheet)
	err = file.Save("./sample.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func addStateNameLine(s *xlsx.Sheet) {
	row := s.Row(0)
	cell := row.AddCell()
	cell.Value = "Event \\ State"
	for i, v := range stateMachineInfo.States {
		cell = row.AddCell()
		cell.SetValue(strconv.Itoa(i) + ":" + v)
	}

}
