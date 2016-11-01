package util

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"os"
	"runtime"
	"encoding/csv"
)

func ExportFile(result *ResultDataSchema, extensions string,columnSchema []ColumnSchema)(string,error)  {
	switch extensions {
		case "xlsx":
			return  exportExcel(result,extensions,columnSchema)
		case "csv":
			return exportCSV(result,extensions,columnSchema)
		default:
			return exportExcel(result,extensions,columnSchema)
	}
}

func getFilePath(extensions string) string  {
	pwd, _ := os.Getwd()
	var filePath string
	if runtime.GOOS == "windows" {
		filePath = fmt.Sprintf("%s\\static\\tmpFile\\mbData.%s",pwd,extensions)
	}else {
		filePath = fmt.Sprintf("%s/static/tmpFile/mbData.%s",pwd,extensions)
	}
	return  filePath
}

//export xlsx file
func exportExcel(result *ResultDataSchema,extensions string,columnSchema []ColumnSchema) (string,error)  {
	var file  *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	
	file = xlsx.NewFile()
	sheet , err = file.AddSheet("mbDataSheet")
	if err != nil{
		fmt.Print(err)
	}
	//add first row for column header in excel
	row = sheet.AddRow()
	colLenght := len(result.Columns)
	for i :=0; i< colLenght; i++ {
		cell = row.AddCell()
		colText := getColumnTitle(result.Columns[i].Text,columnSchema )
		cell.Value = colText
	}
	//end
	
	//add row data
	rowLength := len(result.Rows)
	for j:=0; j< rowLength; j++{
		row = sheet.AddRow()
		for k :=0; k< colLenght; k++ {
			cell = row.AddCell()
			str:= result.Rows[j][k]
			switch str.(type) {
				case float64:
					cell.SetFloat(str.(float64))
				default:
					cell.Value = fmt.Sprintf("%v",str)
				}
			
		}
	}
	//end
	
	filePath :=getFilePath(extensions)
	fmt.Print(filePath)
	err = file.Save(filePath)
	if err != nil{
		fmt.Print(err)
	}
	return  filePath, err
}

//export csv file
func exportCSV(result *ResultDataSchema,extensions string,columnSchema []ColumnSchema) (string,error) {
	filePath :=getFilePath(extensions)
	file, err := os.Create(filePath)
	if err != nil{
		fmt.Print(err)
	}
	defer file.Close()
	
	file.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	
	//write header column
	writer := csv.NewWriter(file)
	
	cols := []string{}
	for _,col := range result.Columns{
		colText := getColumnTitle(col.Text,columnSchema )
		cols = append(cols,colText)
	}
	writer.Write(cols)
	//end
	
	rows := [][]string{}
	for _, row := range result.Rows{
		copyRow :=[]string{}
		for k :=0; k< len(result.Columns); k++ {
			str := row[k]
			var colValue string
			switch str.(type) {
				case float64:
					colValue = fmt.Sprintf("%f",str)
				case float32:
					colValue = fmt.Sprintf("%f",str)
				default:
					colValue = fmt.Sprintf("%v",str)
				}
			copyRow = append(copyRow,colValue)
		}
		rows = append(rows,copyRow)
	}
	writer.WriteAll(rows)
	rows = nil // clear rows object
	defer writer.Flush()
	return filePath,err
}

func getColumnTitle(colField string,columnSchema []ColumnSchema) string  {
	for _, column := range columnSchema {
		if column.Field == colField{
			return column.Name
		}
	}
	return ""
}
