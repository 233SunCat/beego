package controllers

import (
	"beeTest01/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"strings"
)

type StartController struct {
	beego.Controller
}
type article struct {
	Title string
}

func (c *StartController) Get() {
	o := orm.NewOrm()
	txt := []models.Txt{}
	_, err := o.QueryTable("txt").OrderBy("-id").All(&txt)
	if err != nil {
		beego.Info("查询失败", err)
	}
	vs := []models.Txt{}
	for _, t := range txt {
		td := models.Txt{}
		td.Id = t.Id
		td.Title = t.Title
		td.Type = t.Type
		td.CreateAt = t.CreateAt
		len := len(t.Content)
		content := ""
		if len > 50 {
			content = t.Content[:50]
		} else {
			content = t.Content[:len/2]
		}
		td.Content = content
		vs = append(vs, td)
	}
	c.Data["article"] = vs //想办法只获取文本内容前五十文字
	c.TplName = "start.html"

}
