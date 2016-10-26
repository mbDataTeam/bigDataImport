package util

import (
	"fmt"
	"io/ioutil"
	"strings"
	"net/http"
	"encoding/json"
	"bigDataImport/setting"
)

func QueryData(sql string) *ResultDataSchema  {
	url := setting.DataQuery
	fmt.Print(sql)
	
	bodyType :="text/plain" // application/x-www-form-urlencoded
	b := strings.NewReader(sql)
	resp,err := http.Post(url,bodyType,b)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Println(err)
		return &ResultDataSchema{}
	}
	var rowResults []ResultDataSchema
	json.Unmarshal(body,&rowResults)
	if len(rowResults) > 0 {
		return &rowResults[0]
	}
	return &ResultDataSchema{}
}
