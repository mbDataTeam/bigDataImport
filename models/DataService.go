package models

import (
	"fmt"
	"bigDataImport/util"
)

func GetDataList(pageIndex,start,pageCount int,tableName,filters string) (interface{},*util.ResultDataSchema) {
	rows, result := generateRows(pageIndex,start,pageCount,tableName,filters)
	jsonData := map[string]interface{}{}
	totalCount := len(rows)
	jsonData["data"] = rows
	jsonData["rowCount"] = totalCount
	return jsonData, result
}

func generateRows(pageIndex,start,pageCount int, tableName string, filters string) ([]interface{},*util.ResultDataSchema) {
	rowResult := getRows(pageIndex,start,pageCount,tableName,filters)
	colCount := len(rowResult.Columns)
	rowCount := len(rowResult.Rows)
	rows := []interface{}{}
	if rowCount > 0{
		for i:= 0; i < rowCount;i++{
			row := make(map[string]interface{})
			rowData := rowResult.Rows[i]
			for j :=0; j < colCount; j++ {
				k := rowResult.Columns[j].Text
				row[k] = rowData[j]
			}
			rows = append(rows,row)
		}
	}
	//fmt.Print(rows)
	return rows,rowResult;

}

func getRows(pageIndex,start,pageCount int,tableName, filters string) *util.ResultDataSchema {
	limit := 1000
	var sql string
	if filters == ""{
		sql = fmt.Sprintf("SELECT * FROM %s LIMIT %d ",tableName,limit)
	}else {
		sql = fmt.Sprintf("SELECT * FROM %s where %s LIMIT %d", tableName, filters, limit)
	}
	return  util.QueryData(sql)
}
