package util

import (
	"encoding/json"
)

func GenerateFilterAndGridColumns(schema TableSchema)(string, string,[]string,string ){
	total := len(schema.Columns)
	columns := []interface{}{}
	filters := []Filters{}
	fields := []string{}
	if total == 0 {
		return "", "", fields,""
	}
	for i:= 0; i<total;i++ {
		if(schema.Columns[i].Show == true) {
			columns = buildGridColumnsData(schema.Columns, i, columns)
			filters = buildFilterData(schema.Columns, i, filters);
			fields = append(fields, schema.Columns[i].Field)
		}
	}
	columnByte, _ :=json.Marshal(columns)
	filerByte,_ :=json.Marshal(filters)
	
	return string(columnByte),string(filerByte),fields, schema.TableName
}

// build filter schema data
func buildFilterData(columns []ColumnSchema, index int, filters []Filters) []Filters {
	
	filters = append(filters, Filters{
		Id: columns[index].Field,
		Label: columns[index].Name,
		Type: columns[index].Type,
		Input: buildInput(columns[index].Type, columns[index].Values),
		Values: buildValues(columns[index].Values),
		Operators: buildOperations(columns[index].Type),
		Plugin: buildPlug(columns[index].Type),
		Plugin_config: buildPlugConfig(columns[index].Type),
	})
	return filters
}

//build plugin name
func buildPlug(dataType string) (string)  {
	if dataType == Col_Time || dataType == Col_DateTime || dataType == Col_Date{
		return "datepicker"
	}
	return ""
}

//build plugin config
func buildPlugConfig(dataType string) (interface{})  {
	var config interface{}
	switch dataType {
		case Col_Date:
			config = map[string]interface{}{
				"format": "yyyy-mm-dd",
				"todayBtn": "linked",
				"todayHighlight": true,
				"autoclose": false,
			}
		case Col_DateTime:
			config = map[string]interface{}{
				"format": "yyyy-mm-dd",
				"todayBtn": "linked",
				"todayHighlight": true,
				"autoclose": false,
			}
		case Col_Time:
			config = map[string]interface{}{
				"format": "hh:ii:ss",
				"autoclose": false,
			}
	}
	
	return config
}

//return input control
func buildInput(dataType,valueType string) (string) {
	var input string
	switch dataType {
		case Col_Int, Col_Double:
			if valueType == Confirm_Type {
				input = Input_radio
			}else {
				input = Input_text
			}
		case Col_String, Col_Date, Col_DateTime,Col_Time:
			input = Input_text
		default:
			input = Input_text
		}
		return input
}

//return operations
func buildOperations(dataType string) ([]string) {
	var operations []string
		switch dataType {
						
		case Col_Int, Col_Double :
			operations =[]string{ Opt_equal, Opt_not_equal,Opt_in,Opt_not_in,Opt_less,Opt_less_or_equal,Opt_greater,
				Opt_greater_or_equal,Opt_between,Opt_not_between }
			
		case Col_Date, Col_DateTime, Col_Time:
			operations = []string{ Opt_greater, Opt_greater_or_equal,Opt_between,Opt_not_between }
			
		case Col_String:
			operations =[]string{ Opt_equal, Opt_not_equal,Opt_in,Opt_not_in,Opt_begins_with,Opt_not_begins_with,
				Opt_contains, Opt_not_contains,Opt_ends_with,Opt_not_ends_with,Opt_is_empty,Opt_is_not_empty }
			}

	return operations
}

//return values
func buildValues(valuesType string) (interface{}) {
	var values interface{}
	switch valuesType {
		case Confirm_Type:
			values = InitYesOrNo()
		case Gender_Type:
			values = InitGenders()
		default:
			values = ""
		}
	return values
}

//build grid column schema data
func buildGridColumnsData(columns []ColumnSchema ,index int, cols []interface{})[]interface{}{
	cols = append(cols, ColumnDefinition{
		Text: columns[index].Name,
		DataIndex: columns[index].Field,
		Width : 120,
		Sortable:true,
		DataType: columns[index].Type,
	})
	return cols;
}
