package models

import (
	"fmt"
	"bigDataImport/util"
	"bigDataImport/setting"
	"strings"
)

func GetDataList(pageIndex,start,pageCount int,tableName,filters,companyId string) (interface{},*util.ResultDataSchema) {
	rows, result := generateRows(pageIndex,start,pageCount,tableName,filters,companyId)
	jsonData := map[string]interface{}{}
	totalCount := len(rows)
	jsonData["data"] = rows
	jsonData["rowCount"] = totalCount
	return jsonData, result
}

func generateRows(pageIndex,start,pageCount int, tableName, filters,companyId string) ([]interface{},*util.ResultDataSchema) {
	rowResult := getRows(pageIndex,start,pageCount,tableName,filters,companyId)
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

func getRows(pageIndex,start,pageCount int,tableName, filters,companyId string) *util.ResultDataSchema {
	find :=strings.Index(tableName,"?")
	if find >=0{
		tableName = strings.Replace(tableName,"?",companyId,-1)
	}
	limit := setting.Limit
	var sql string
	if filters == ""{
		sql = fmt.Sprintf("SELECT * FROM %s LIMIT %d ",tableName,limit)
	}else {
		sql = fmt.Sprintf("SELECT * FROM %s where %s LIMIT %d", tableName, filters, limit)
	}
	return  util.QueryData(sql)
}

//如果查询条件里面有 更新日期 需要替换 格式：20160809
func processPartionDate(filters, tableName string) string{
	 if filters == ""{
		 return ""
	 }
	 updateTimeKey := "user_task_update_time"
	 findKey := strings.Index(filters,updateTimeKey)
	 if findKey >=0{
		 
	 }
	return ""
}
