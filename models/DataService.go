package models

func GetDataList(pageIndex,start,pageCount int) interface{}{
	rows := generateRows(pageIndex,start,pageCount)
	jsonData := map[string]interface{}{}
	totalCount := 8;
	jsonData["data"] = rows
	jsonData["rowCount"] = totalCount

	return jsonData
}

func generateRows(pageIndex,start,pageCount int) interface{} {
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
