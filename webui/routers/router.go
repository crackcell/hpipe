package routers

import (
	"github.com/crackcell/hpipe/webui/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
