package models

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bigDataImport/util"
)

func GetDataList(pageIndex,start,pageCount int,tableName string) interface{}{
	rows := generateRows(pageIndex,start,pageCount,tableName)
	jsonData := map[string]interface{}{}
	totalCount := len(rows)
	jsonData["data"] = rows
	jsonData["rowCount"] = totalCount

	return jsonData
}

func generateRows(pageIndex,start,pageCount int, tableName string) []interface{} {
	rowResult := getRows(pageIndex,start,pageCount,tableName)
	colCount := len(rowResult.Columns)
	rowCount := len(rowResult.Rows)
	rows := []interface{}{}
	if rowCount > 0{
		for i:= 0; i < rowCount;i++{
			row := make(map[string]interface{})
			rowData := rowResult.Rows[i]
			for j :=0; j < colCount; j++ {
				k := rowResult.Columns[j].Text
				row[k] = rowData[j]  //TODO if column is datetime or data, should format
			}
			rows = append(rows,row)
		}
	}
	//fmt.Print(rows)
	return rows;

}

func getRows(pageIndex,start,pageCount int,tableName string) *util.ResultDataSchema {
	url := "http://192.168.174.135:8085/query"
	limit := 10000
	postData := fmt.Sprintf("SELECT * FROM %s LIMIT %d",tableName,limit)
	bodyType :="text/plain" //application/x-www-form-urlencoded
 	b := strings.NewReader(postData)
	resp,err := http.Post(url,bodyType,b)
	if err != nil {
		fmt.Println(err)
		return &util.ResultDataSchema{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var rowResults []util.ResultDataSchema
	json.Unmarshal(body,&rowResults)
	fmt.Println(string(body))
	if len(rowResults) > 0 {
		return &rowResults[0]
	}
	return &util.ResultDataSchema{}
}
