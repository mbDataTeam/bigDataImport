package models

import (
	"fmt"
	"bigDataImport/util"
	"bigDataImport/setting"
	"strings"
	"strconv"
)

func GetDataList(pageIndex,start,pageCount int,tableName,sqlView,filters,companyId string) (interface{},*util.ResultDataSchema) {
	rows, result := generateRows(pageIndex,start,pageCount,tableName,sqlView,filters,companyId)
	jsonData := map[string]interface{}{}
	totalCount := len(rows)
	jsonData["data"] = rows
	jsonData["rowCount"] = totalCount
	return jsonData, result
}

func generateRows(pageIndex,start,pageCount int, tableName,sqlView, filters,companyId string) ([]interface{},*util.ResultDataSchema) {
	rowResult := getRows(pageIndex,start,pageCount,tableName, sqlView, filters,companyId)
	colCount := len(rowResult.Columns)
	rowCount := len(rowResult.Rows)
	rows := []interface{}{}
	if rowCount > 0{
		for i:= 0; i < rowCount;i++{
			if i == setting.Top {  // setting grid data count
				break;
			}
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

func getRows(pageIndex,start,pageCount int,tableName, sqlView, filters,companyId string) *util.ResultDataSchema {
	find :=strings.Index(sqlView,"?")
	if find >=0{
		sqlView = strings.Replace(sqlView,"?",companyId,-1)
	}
	limit := setting.Limit
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

