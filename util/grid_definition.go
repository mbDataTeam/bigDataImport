package util

type ResultDataSchema struct {
	Cols []Columns `json:"columns"`
	Rows [][]interface{} `json:"rows"`
	Type string `json:"type"`
}

type Columns struct {
	Text string `json:"text"`
}

type TableSchema struct {
	MetaId string `json:"meta_id"`
	TableDesc string `json:"table_desc"`
	SelectGroup string `json:"select_group"` // 确定是否显示一组级联dropdownlist
	TableName string `json:"table_name"`
	Columns []ColumnSchema `json:"columns"`
}

type ColumnSchema struct {
	Name string `json:"name"`
	Field string `json:"field"`
	Type string `json:"type"`
	Show bool `json:"show"`
	Values string `json:"values"`
	ReferField string `json:"refer_field"`
	FilterShow bool `json:"filter_show"`
}

type Filters struct {
	Id string `json:"id"`
	Label string `json:"label"`
	Type string `json:"type"`     //string, integer, double, date, time, datetime and boolean.
	Input string `json:"input"`  //text, textarea, radio, checkbox and select
	Values interface{} `json:"values"` // catalog data or checkbox data
	Operators []string `json:"operators"` //
	Plugin string `json:"plugin"`
	Plugin_config interface{} `json:"plugin_config"`
}

type ColumnDefinition struct {
	Text string              `json:"text"`
	Width int32              `json:"width"`
	DataIndex string         `json:"dataIndex"`
	DataType string          `json:"data_type"`
	Sortable bool            `json:"sortable"`
}

type ImportDataDefinition struct{
	GridTitle string
	Columns string
	Filters string
	Fields []string
	TableName string
	CompanyId string
	SelectGroup string
}

type SelectSchema struct {
	GroupName string
	OptionText string
	OptionValue string
}

//enum column data type
const(
	Col_Int   = "integer"
	Col_Double = "double"
	Col_String = "string"
	Col_Time  =   "time"
	Col_Date   = "date"
	Col_DateTime = "datetime"
	//Col_Boolean = "boolean"
)

//filter operations enum
const (
    Opt_equal = "equal"                          //apply_to: ['string', 'number', 'datetime', 'boolean']
    Opt_not_equal = "not_equal"                 //apply_to: ['string', 'number', 'datetime', 'boolean']
	Opt_in = "in"                                 //apply_to: ['string', 'number', 'datetime']
	Opt_not_in = "not_in"                        //apply_to: ['string', 'number', 'datetime']
	Opt_less ="less"                              //apply_to: ['number', 'datetime']
	Opt_less_or_equal = "less_or_equal"        //apply_to: ['number', 'datetime']
	Opt_greater = "greater"                     //['number', 'datetime']
	Opt_greater_or_equal ="greater_or_equal" //['number', 'datetime']
	Opt_between ="between"                     //['number', 'datetime']
	Opt_not_between = "not_between"           //['number', 'datetime']
	Opt_begins_with ="begins_with"            //['string']
	Opt_not_begins_with="not_begins_with"    //['string']
	Opt_contains = "contains"                  //['string']
	Opt_not_contains ="not_contains"          //['string']
	Opt_ends_with = "ends_with"                //['string']
	Opt_not_ends_with = "not_ends_with"       //['string']
	Opt_is_empty = "is_empty"                   //['string']
	Opt_is_not_empty = "is_not_empty"          //['string']
	Opt_is_null = "is_null"                     // ['string', 'number', 'datetime', 'boolean']
	Opt_is_not_null = "is_not_null"            // ['string', 'number', 'datetime', 'boolean']
 )

//filter input enum
const (
	Input_radio = "radio"
	Input_checkbox = "checkbox"
	Input_select = "select"
	Input_text = "text"
	Input_textarea = "textarea"
)

const (
	Confirm_Type = "confirm" //是 否
	Gender_Type = "gender" //男 女
	Conditon_Type = "conditionType" //条件类型
	Task_Type = "taskType" //任务类型
	Emp_Status_Type = "empStatus" //员工状态
)

const (
	Sel_ParentCatalog = "pCatalog"
	Sel_Catalog = "catalog"
	Sel_Course = "course"
)

const (
	Course_Group = "course_group"
)
