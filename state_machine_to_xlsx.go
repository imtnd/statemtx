package main

import (
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"
)

func makeXlsx() {
	fmt.Println(stateMachineInfo)

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	addStateNameLine(sheet)
	addEventNameLine(sheet)
	err = file.Save(outputFileName)
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

func addEventNameLine(s *xlsx.Sheet) {
	for i, v := range stateMachineInfo.Events {
		row := s.Row(1 + i)
		cell := row.AddCell()
		cell.Value = v
	}
}
