package util

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"os"
	"runtime"
)

//export xlsx file
func ExportExcel() (string,error)  {
	var file  *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	
	file = xlsx.NewFile()
	sheet , err = file.AddSheet("Shee1")
	if err != nil{
		fmt.Print(err)
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "test xlsx"
	
	pwd, _ := os.Getwd()
	var filePath string
	if runtime.GOOS == "windows" {
		filePath = pwd + "\\static\\tmpFile\\mbData.xlsx"
	}else {
		filePath = pwd + "/static/tmpFile/mbData.xlsx"
	}
	fmt.Print(filePath)
	err = file.Save(filePath)
	if err != nil{
		fmt.Print(err)
	}
	return  filePath,err
}

//export csv file
func ExportCSV()  {
	
}
