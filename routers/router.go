package routers

import (
	"bigDataImport/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/api/import", &controllers.ImportController{})
    beego.Router("/api/fetchData", &controllers.ImportController{},"*:List")
}
