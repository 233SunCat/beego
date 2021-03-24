package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/index.tpl"
}
