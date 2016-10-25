package routers

import (
	"bigDataImport/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/api/import", &controllers.ImportController{})
    beego.Router("/api/fetchData", &controllers.ImportController{},"*:List")
	
	beego.Router("/api/importData", &controllers.ImportController{},"Post:SaveFile")
	
	beego.Router("/api/fillSelect", &controllers.ImportController{} , "Post:FillDropdownData")
}
