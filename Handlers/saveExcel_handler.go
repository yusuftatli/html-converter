package handlers

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/orgs/mdyazilim/html-converter-new/common"
	"github.com/orgs/mdyazilim/html-converter-new/models"
)

func SaveToExcel(orders[] models.ExcelDto, fileName string, _packageCreataDate string, _packageNumber string) {

	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")

			count := 1
	for _,
	element := range orders {

			f.SetCellValue("Sheet1", "A" + strconv.Itoa(count), element.Col18)
			f.SetCellValue("Sheet1", "B" + strconv.Itoa(count), element.Col0)
			f.SetCellValue("Sheet1", "C" + strconv.Itoa(count), element.Col1)
			f.SetCellValue("Sheet1", "D" + strconv.Itoa(count), element.Col2)
			f.SetCellValue("Sheet1", "E" + strconv.Itoa(count), element.Col3)
			f.SetCellValue("Sheet1", "F" + strconv.Itoa(count), element.Col4)
			f.SetCellValue("Sheet1", "G" + strconv.Itoa(count), element.Col5)
			f.SetCellValue("Sheet1", "H" + strconv.Itoa(count), element.Col6)
			f.SetCellValue("Sheet1", "I" + strconv.Itoa(count), element.Col7)
			f.SetCellValue("Sheet1", "J" + strconv.Itoa(count), element.Col8)
			f.SetCellValue("Sheet1", "K" + strconv.Itoa(count), element.Col9)
			f.SetCellValue("Sheet1", "L" + strconv.Itoa(count), element.Col10)
			f.SetCellValue("Sheet1", "M" + strconv.Itoa(count), element.Col11)
			f.SetCellValue("Sheet1", "N" + strconv.Itoa(count), element.Col12)
			f.SetCellValue("Sheet1", "O" + strconv.Itoa(count), element.Col13)
			f.SetCellValue("Sheet1", "P" + strconv.Itoa(count), element.Col14)
			f.SetCellValue("Sheet1", "Q" + strconv.Itoa(count), element.Col15)
			f.SetCellValue("Sheet1", "R" + strconv.Itoa(count), element.Col16)
			f.SetCellValue("Sheet1", "S" + strconv.Itoa(count), element.Col17)
			f.SetCellValue("Sheet1", "S" + strconv.Itoa(count), element.Col17)
			f.SetCellValue("Sheet1", "T" + strconv.Itoa(count), _packageCreataDate)
			f.SetCellValue("Sheet1", "U" + strconv.Itoa(count), _packageNumber)
			count++
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs(common.GetEnvironment().MainRoot + "exels/" + fileName + ".xlsx");err != nil {
			fmt.Println(err)
	}
}