package models

import (
	"fmt"
	"bigDataImport/util"
	"strings"
	"strconv"
)

func GetDataList(tableName,sqlView,filters,companyId string,limit int) (interface{},[]interface{},[]util.Columns) {
	rows,columns := generateRows(tableName,sqlView,filters,companyId,limit)
	jsonData := map[string]interface{}{}
	totalCount := len(rows)
	jsonData["data"] = rows
	jsonData["rowCount"] = totalCount
	return jsonData,rows,columns
}

func generateRows(tableName,sqlView, filters,companyId string,limit int) ([]interface{},[]util.Columns) {
	rowResult := GetRows(tableName, sqlView, filters,companyId,limit)
	colCount := len(rowResult.Cols)
	rowCount := len(rowResult.Rows)
	rows := []interface{}{}
	if rowCount > 0{
		for i:= 0; i < rowCount;i++{
			row := make(map[string]interface{})
			rowData := rowResult.Rows[i]
			for j :=0; j < colCount; j++ {
				k := rowResult.Cols[j].Text
				row[k] = rowData[j]
			}
			rows = append(rows,row)
		}
	}
	//fmt.Print(rows)
	return rows,rowResult.Cols;

}

func GetRows(tableName, sqlView, filters,companyId string, limit int) *util.ResultDataSchema {
	find :=strings.Index(sqlView,"?")
	if find >=0{
		sqlView = strings.Replace(sqlView,"?",companyId,-1)
	}
	var sql string
	switch tableName {
		case "task_data_export_view":
			sql = `select * from (select a.* ,row_number() over(partition by a.task_user_id order by task_update_time desc)
											rank         from ( `+ sqlView +` ) a where
											`+ filters +` ) where rank=1 LIMIT ` + strconv.Itoa(limit)
		default:
			sql = fmt.Sprintf(`SELECT * FROM %s where company_id in ( %s ) and %s LIMIT %d`, sqlView, companyId, filters, limit)
	}
	
	return  util.QueryData(sql)
}

