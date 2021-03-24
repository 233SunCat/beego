package controllers

import (
	"beeTest01/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//"github.com/astaxie/beego/orm"
	//"github.com/hzde0128/hblog/models"
	//"net/http"
	"strconv"
	//"strings"
)

type BlogController struct {
	beego.Controller
}

// 文章列表页面
func (c *BlogController) Get() {
	o := orm.NewOrm()
	txt := []models.Txt{}
	_, err := o.QueryTable("txt").OrderBy("-id").All(&txt)
	if err != nil {
		beego.Info("查询失败", err)
	}
	c.Data["article"] = txt
	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/list.tpl"
}

// 添加文章
func (c *BlogController) Add() {
	if c.Ctx.Request.Method == "POST" {
		content := c.GetString("content")
		content = strings.Replace(content, "\n", "<br>", -1)
		article := models.Txt{
			Title:   c.GetString("title"),
			Type:    c.GetString("type"),
			Content: content,
		}
		beego.Info("添加文章内容:", article)
		article.Insert(&article)
		article.ReFresh() //刷新id
	}

	//添加文章内容
	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/article.tpl"
}

//删除文章
//a标签属于get请求
func (c *BlogController) Delete() {
	//o := orm.NewOrm()
	tt := models.Txt{}
	/* 	if c.Ctx.Request.Method == "POST" {
		beego.Info("正确的post响应\n")
		beego.Info(c.Ctx.Request.Host) //127.0.0.1:8080
		beego.Info(c.Ctx.Request.Form) //map[city:[qqqqqqqqqq] id:[10] name:[老鼠爱大米]]
		beego.Info(c.Ctx.Request.Form["name"])
	} */
	if c.Ctx.Request.Method == "POST" {
		str := c.Ctx.Request.Form["id"]
		id, _ := strconv.Atoi(str[0])
		tt.Id = id
		tt.Delete()
	} else {
		//tt.Id = 10
		beego.Info("报错:++++++++++++++")
		//tt.Delete()
	}
	tt.ReFresh()
	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/list.tpl"
}

var a []int = []int{}

//编辑//暂时使用添加更改一下//获取文章的标题以及分类
func (c *BlogController) GetID() {

	//c.Redirect("/admin/blog/edit", http.StatusFound)
	if c.Ctx.Request.Method == "POST" {
		str := c.Ctx.Request.Form["id"]
		id, _ := strconv.Atoi(str[0])
		a = append(a, id)
	}
	beego.Info("数组ID: ", a)
	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/article2.tpl"
}
func (c *BlogController) Edit() {
	id := -1

	if len(a) != 0 {
		id = a[0]

		o := orm.NewOrm()
		tt := models.Txt{}
		tt.Id = id
		err := o.Read(&tt)
		if err != nil {
			beego.Info("查询失败", err)
			return
		}
		beego.Info("查询成功", tt)

		c.Data["title"] = tt.Title
		c.Data["type"] = tt.Type
		content := strings.Replace(tt.Content, "<br>", "\n", -1)
		c.Data["content"] = content

	}
	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/article2.tpl"
}
func (c *BlogController) Update() {
	if c.Ctx.Request.Method == "POST" {
		content := c.GetString("content")
		content = strings.Replace(content, "\n", "<br>", -1)
		article := models.Txt{
			Title:   c.GetString("title"),
			Type:    c.GetString("type"),
			Content: content,
		}
		article.Id = a[0]
		beego.Info("修改文章内容:", article)
		article.Update(&article)
	}
	a = []int{}
	//修改文章内容
	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/article2.tpl"
}
