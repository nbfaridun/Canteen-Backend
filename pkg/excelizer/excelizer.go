package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// Create a new Excel file
	file := excelize.NewFile()

	// Create a new sheet
	sheetName := "Sheet1"
	//index := file.NewSheet(sheetName)

	// Set the column headers and apply formatting
	headers := map[string]string{
		"A1": "Food",
		"B1": "Students",
		"C1": "Cost",
		"D1": "Time",
	}

	for cell, value := range headers {
		file.SetCellValue(sheetName, cell, value)
		styleID, _ := file.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal":"center"},"border":[{"type":"bottom","color":"000000","style":2}],"fill":{"type":"pattern","color":["F2F2F2"],"pattern":1}}`)
		file.SetCellStyle(sheetName, cell, cell, styleID)
	}

	// Add some sample data and apply formatting
	data := map[string]interface{}{
		"A2": "Pizza",
		"B2": "John Doe",
		"C2": 10.5,
		"D2": "12:30 PM",
	}

	for cell, value := range data {
		file.SetCellValue(sheetName, cell, value)
		styleID, _ := file.NewStyle(`{"alignment":{"horizontal":"center"},"border":[{"type":"bottom","color":"000000","style":1}]}`)
		file.SetCellStyle(sheetName, cell, cell, styleID)
	}

	// Increase the font size for all cells
	file.SetColWidth(sheetName, "A", "D", 15)

	// Save the file with a given name
	fileName := "beautiful_example.xlsx"
	if err := file.SaveAs(fileName); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Excel file '%s' created successfully.\n", fileName)
}
