package models

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bigDataImport/util"
)

func GetDataList(pageIndex,start,pageCount int) interface{}{
	rows := generateRows(pageIndex,start,pageCount)
	jsonData := map[string]interface{}{}
	totalCount := 8;
	jsonData["data"] = rows
	jsonData["rowCount"] = totalCount

	return jsonData
}

func generateRows(pageIndex,start,pageCount int) interface{} {
	rowResult := getRows()
	
	fmt.Print(rowResult)
	rows := []interface{}{}
	if pageIndex > 1{
		row5 := map[string]interface{}{
			"Id": 5,
			"Author":"bill5",
			"Title": "title difinition5",
			"Manufacturer":"MB5",
			"Product":"course5",
		}
		rows = append(rows,row5)
		row6 := map[string]interface{}{
			"Id": 6,
			"Author":"bill6",
			"Title": "title difinition6",
			"Manufacturer":"MB6",
			"Product":"course6",
		}
		rows = append(rows,row6)
		row7 := map[string]interface{}{
			"Id": 7,
			"Author":"bill7",
			"Title": "title difinition7",
			"Manufacturer":"MB7",
			"Product":"course7",
		}
		rows = append(rows,row7)

		row8 := map[string]interface{}{
			"Id": 8,
			"Author":"bill8",
			"Title": "title difinition8",
			"Manufacturer":"MB8",
			"Product":"course8",
		}
		rows = append(rows,row8)
	}else{
		row1 := map[string]interface{}{
		"Id": 1,
		"Author":"bill1",
		"Title": "title difinition1",
		"Manufacturer":"MB1",
		"Product":"course1",
	}
		rows = append(rows,row1)
		row2 := map[string]interface{}{
			"Id": 2,
			"Author":"bill2",
			"Title": "title difinition2",
			"Manufacturer":"MB2",
			"Product":"course2",
		}
		rows = append(rows,row2)
		row3 := map[string]interface{}{
			"Id": 3,
			"Author":"bill3",
			"Title": "title difinition3",
			"Manufacturer":"MB3",
			"Product":"course3",
		}
		rows = append(rows,row3)
		row4 := map[string]interface{}{
			"Id": 4,
			"Author":"bill4",
			"Title": "title difinition4",
			"Manufacturer":"MB4",
			"Product":"course4",
		}
		rows = append(rows,row4)
	}
	return rows;

}

func getRows() util.ResultDataSchema {
	url := "http://192.168.174.135:8085/query"
	postData := `SELECT * FROM mb.warehouse.course_feedback_raw LIMIT 1000`
	bodyType :="text/plain" //application/x-www-form-urlencoded
 	b := strings.NewReader(postData)
	resp,err := http.Post(url,bodyType,b)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var rowResults []util.ResultDataSchema
	json.Unmarshal(body,&rowResults)
	fmt.Println(string(body))
	return rowResults[0]
}
