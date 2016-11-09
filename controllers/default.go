package controllers

import (
	"github.com/astaxie/beego"
	"bigDataImport/util"
	"strconv"
	"bigDataImport/models"
	"fmt"
	"bigDataImport/setting"
	"encoding/json"
)

var (
	bigDataResult *util.ResultDataSchema
	viewName string
	companyIds string
	tbName string
)

type ImportController struct {
	beego.Controller
}

// http://localhost:9100/api/import/?meta_id=task_data_export&company_id=207
// http://localhost:9100/api/import/?meta_id=task_data_export&company_id=207&sign=6fc0b7963bf543a59b506430475f204a324a2479
func (c *ImportController) Get() {
	
	/*var getUrl string
	requestUrl,_ :=  url.QueryUnescape(c.Ctx.Request.URL.String())
	getUrl = strings.Join([]string{setting.SignUrl,requestUrl},""); //http://databi.ifuli.cn:39200
	if (util.ValidateSignUrl(getUrl) == false){
		c.Ctx.WriteString("invalidate sign name")
	}*/
	
	metaId := c.GetString("meta_id") // get table name
	companyIds = c.GetString("company_id")
	tableSchema := util.QueryTableMeta(metaId); // query table schema from elastic search
	//tableSchema.SelectGroup = "Course_Group"; // todo remove
	jsonColumn, jsonFilter, fields,tableName := util.GenerateFilterAndGridColumns(*tableSchema)
	//columnSchema = tableSchema.Columns
	viewName = setting.SQLView[tableName]
	tbName = tableName
	
	c.Data["ImportDataDefinition"] = &util.ImportDataDefinition{
		GridTitle: tableSchema.TableDesc,
		Columns: jsonColumn,
		Filters: jsonFilter,
		Fields:  fields,
		SelectGroup: tableSchema.SelectGroup,
	}
	c.TplName = "import.tpl"
}

//paging
func (c *ImportController) List() {
	page,_ := strconv.Atoi(c.GetString("page"))   // page index, start with 1
	start,_ := strconv.Atoi(c.GetString("start"))  // start row index , start with 0
	limit,_ := strconv.Atoi(c.GetString("limit"))  // row count per page
	//tableName := c.GetString("tableName")      // table name
	filters := c.GetString("filters")          // sql where condition
	jsonData,result:= models.GetDataList(page,start,limit,tbName,viewName,filters,companyIds)
	bigDataResult = result
	c.Data["json"] = jsonData
	c.ServeJSON()
}

//generate csv or excel file
func (c *ImportController) SaveFile(){
	extensions := c.GetString("extensions")      // file extensions name
	cols := c.GetString("cols")
	schema := make([]util.ColumnSchema,0)
	json.Unmarshal([]byte(cols),&schema)
	fmt.Sprintf("download schema %s", schema)
	_,err := util.ExportFile(bigDataResult,extensions,schema)
	var jsonData string
	if err == nil{
		jsonData = `{"successful" : true }`
	}else {
		jsonData = `{"successful" : false }`
	}
	c.Data["json"] = jsonData
	c.ServeJSON()
}

func (c *ImportController) FillDropdownData()  {
	ids := c.GetString("ids")
	selType := c.GetString("selectType")
	//compIds := c.GetString("compIds")
	//var jsonData []util.SelectSchema
	switch selType {
		case util.Sel_ParentCatalog:
			c.Data["json"] = util.GetParentCatalogs(companyIds)
		case util.Sel_Catalog:
			c.Data["json"] = util.GetCatlogs(companyIds,ids)
		case util.Sel_Course:
			c.Data["json"] = util.GetCourseList(companyIds,ids)
	}
	fmt.Sprintf("%s",ids)
	//c.Data["json"] = jsonData
	c.ServeJSON()
	
}
