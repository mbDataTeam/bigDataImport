package controllers

import (
	"github.com/astaxie/beego"
	"bigDataImport/util"
	"strconv"
	"bigDataImport/models"
	"fmt"
)

var (
	bigDataResult *util.ResultDataSchema
)

type ImportController struct {
	beego.Controller
}

// http://localhost:9100/api/import/?meta_id=course_feedback_raw&company_id=207
// http://localhost:9100/api/import/?meta_id=course_quiz_raw&company_id=207
func (c *ImportController) Get() {
	//requestUrl,_ := url.QueryUnescape(c.Ctx.Request.URL.String())
	/*if (util.ValidateSignUrl(requestUrl) == false){
		c.Ctx.WriteString("invalidate sign name")
	}*/
	metaId := c.GetString("meta_id") // get table name
	companyId := c.GetString("company_id")
	tableSchema := util.QueryTableMeta(metaId); // query table schema from elastic search
	jsonColumn, jsonFilter, fields,tableName := util.GenerateFilterAndGridColumns(*tableSchema)
	
	//fmt.Print(jsonFilter)
	
	c.Data["ImportDataDefinition"] = &util.ImportDataDefinition{
		GridTitle: tableSchema.TableDesc,
		Columns: jsonColumn,
		Filters: jsonFilter,
		Fields:  fields,
		TableName: tableName,
		CompanyId: companyId,
	}
	c.TplName = "import.tpl"
}

//paging
func (c *ImportController) List() {
	page,_ := strconv.Atoi(c.GetString("page"))   // page index, start with 1
	start,_ := strconv.Atoi(c.GetString("start"))  // start row index , start with 0
	limit,_ := strconv.Atoi(c.GetString("limit"))  // row count per page
	tableName := c.GetString("tableName")      // table name
	filters := c.GetString("filters")          // sql where condition
	jsonData,result:= models.GetDataList(page,start,limit,tableName,filters)
	bigDataResult = result
	c.Data["json"] = jsonData
	c.ServeJSON()
}

//generate csv or excel file
func (c *ImportController) SaveFile(){
	extensions := c.GetString("extensions")      // file extensions name
	fmt.Sprintf("download file extensions %s", extensions)
	_,err := util.ExportFile(bigDataResult,extensions)
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
	compIds := c.GetString("compIds")
	//var jsonData []util.SelectSchema
	switch selType {
		case util.Sel_ParentCatalog:
			c.Data["json"] = util.GetParentCatalogs(compIds)
		case util.Sel_Catalog:
			c.Data["json"] = util.GetCatlogs(compIds,ids)
		case util.Sel_Course:
			c.Data["json"] = util.GetCourseList(compIds,ids)
	}
	fmt.Sprintf("%s",ids)
	//c.Data["json"] = jsonData
	c.ServeJSON()
	
}
