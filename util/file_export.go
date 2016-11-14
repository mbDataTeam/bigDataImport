package util

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"os"
	"runtime"
	"encoding/csv"
)

func ExportFile(resultRows [][]interface{}, extensions string,columnSchema []ColumnSchema, columns []Columns)(string,error)  {
	switch extensions {
		case "xlsx":
			return  exportExcel(resultRows,extensions,columnSchema,columns)
		case "csv":
			return exportCSV(resultRows,extensions,columnSchema,columns)
		default:
			return exportExcel(resultRows,extensions,columnSchema,columns)
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
func exportExcel(resultRows [][]interface{},extensions string,columnSchema []ColumnSchema, columns []Columns) (string,error)  {
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
	for _,col := range columnSchema{
		cell = row.AddCell()
		cell.Value = col.Name
	}
	
	for _, resultRow := range resultRows{
		row = sheet.AddRow()
		for _,col := range columnSchema{
			cell = row.AddCell()
			findIndex := getColumnIndex(col.Field,columns)
			str:= resultRow[findIndex]
			switch str.(type) {
				case float64:
					cell.SetFloat(str.(float64))
				default:
					cell.Value = fmt.Sprintf("%v",str)
				}
		}
	}
	/*row = sheet.AddRow()
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
	*/
	
	filePath :=getFilePath(extensions)
	fmt.Print(filePath)
	err = file.Save(filePath)
	if err != nil{
		fmt.Print(err)
	}
	return  filePath, err
}

//export csv file
func exportCSV(resultRows [][]interface{},extensions string,columnSchema []ColumnSchema,columns []Columns) (string,error) {
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
	for _,col := range columnSchema{
		//colText := getColumnTitle(col.Text,columnSchema )
		cols = append(cols,col.Name)
	}
	writer.Write(cols)
	//end
	
	rows := [][]string{}
	for _, row := range resultRows{
		copyRow :=[]string{}
		for _,col := range columnSchema{
			findIndex := getColumnIndex(col.Field,columns)
			str:= row[findIndex]
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
		/*
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
		}*/
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

func getColumnIndex(field string, cols []Columns ) int  {
	var findIndex int
	for index, col := range cols{
		if col.Text == field {
			findIndex = index
			break
		}
	}
	return findIndex
}
