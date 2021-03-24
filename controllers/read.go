package controllers

import (
	"beeTest01/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ReadController struct {
	beego.Controller
}

func (c *ReadController) Get() {
	id, err := strconv.Atoi(c.Ctx.Request.URL.RawQuery[3:])
	if err != nil {
		beego.Info("报错+++++")
	}
	o := orm.NewOrm()
	txt := models.Txt{}
	_ = o.Raw("select * from txt where id=?", id).QueryRow(&txt)
	c.Data["Title"] = txt.Title
	content := txt.Content
	content = strings.Replace(content, "<br>", "\n", -1)
	beego.Info("数据库读取:", content)
	c.Data["Content"] = content
	c.TplName = "read.html"

}
