package controllers

import (
	"github.com/astaxie/beego"
	"bigDataImport/util"
	"strconv"
	"bigDataImport/models"
	"fmt"
)

type ImportController struct {
	beego.Controller
}

// http://localhost:9100/api/import/?meta_id=course_feedback_raw
// http://localhost:9100/api/import/?meta_id=course_quiz_raw
func (c *ImportController) Get() {
	//requestUrl,_ := url.QueryUnescape(c.Ctx.Request.URL.String())
	/*if (util.ValidateSignUrl(requestUrl) == false){
		c.Ctx.WriteString("invalidate sign name")
	}*/
	metaId := c.GetString("meta_id") // get table name
	tableSchema := util.QueryTableMeta(metaId); // query table schema from elastic search
	jsonColumn, jsonFilter, fields := util.GenerateFilterAndGridColumns(*tableSchema)
	
	fmt.Print(jsonFilter)
	
	c.Data["ImportDataDefinition"] = &util.ImportDataDefinition{
		GridTitle: tableSchema.TableDesc,
		Columns: jsonColumn,
		Filters: jsonFilter,
		Fields:  fields,
	}
	c.TplName = "import.tpl"
}

//paging
func (c *ImportController) List() {
	page,_ := strconv.Atoi(c.GetString("page"))   //page index, start with 1
	start,_ := strconv.Atoi(c.GetString("start"))  // start row index , start with 0
	limit,_ := strconv.Atoi(c.GetString("limit"))  // row count per page

	jsonData:= models.GetDataList(page,start,limit)
	c.Data["json"] = jsonData
	c.ServeJSON()
}
