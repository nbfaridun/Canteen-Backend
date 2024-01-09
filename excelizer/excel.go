package excelizer

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"time"
)

func writeToExcelV1() {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "I am a cell!"

	cell = row.AddCell()
	cell.Value = time.Now().String()

	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Excel file created!")
	}

}
