package routers

import (
	"beeTest01/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login.tpl", &controllers.LoginController{}, "get:Get;post:PostStruct") //默认调用get
	beego.Router("/", &controllers.StartController{})                                     //首页路由控制
	beego.Router("/read.html", &controllers.ReadController{})                             //详细路由控制

	beego.Router("/admin/blog/", &controllers.BlogController{}, "get:Get")
	beego.Router("/admin/blog/add", &controllers.BlogController{}, "*:Add")                //增
	beego.Router("/admin/blog/edit", &controllers.BlogController{}, "get:Edit;post:GetID") //改
	beego.Router("/admin/blog/update", &controllers.BlogController{}, "*:Update")          //递交更新数据
	beego.Router("/admin/blog/delete", &controllers.BlogController{}, "*:Delete")          //记住路由比我想象的要复杂一些,不是看网址显示,估计那个只是显示一部分
	beego.Router("/admin/blog/tag", &controllers.BlogController{}, "*:Delete")
}
