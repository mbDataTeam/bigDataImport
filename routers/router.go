package routers

import (
	"bigDataImport/controllers"
	"github.com/astaxie/beego"
	_"bigDataImport/setting"
	"bigDataImport/setting"
)

func init() {
    beego.Router(setting.SecondUrl+"/api/import", &controllers.ImportController{})
    beego.Router(setting.SecondUrl+"/api/fetchData", &controllers.ImportController{},"*:List")
	
	beego.Router(setting.SecondUrl+"/api/importData", &controllers.ImportController{},"Post:SaveFile")
	
	beego.Router(setting.SecondUrl+"/api/fillSelect", &controllers.ImportController{} , "Post:FillDropdownData")
}
