package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "Hpipe WebUI"
	c.Data["Email"] = "tanmenglong@gmail.com"
	c.TplNames = "index.tpl"
}
