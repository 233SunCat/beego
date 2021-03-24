package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	beego.Info("就是这样+++++++++")
	c.TplName = "login.tpl"
}

type User struct {
	Username string
	Password string
}

//解析表单
func (c *LoginController) PostStruct() {
	var user User
	if error := c.ParseForm(&user); error != nil {
		c.Ctx.WriteString("出错了")
	}
	if user.Username == "蒋政" && user.Password == "123456789" {
		beego.Info("就是这样+++++++++")
		c.Redirect("/admin/blog/", http.StatusFound) //跳转到后台
	}
}
