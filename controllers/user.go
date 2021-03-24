package controllers

import "github.com/astaxie/beego"
import "fmt"
//创建控制器的结构体
type UserController struct {
	beego.Controller
}
func (c *UserController)Get(){
	c.Ctx.WriteString("用户中心")
}
//get
func (c *UserController)AddUser(){
	c.TplName = "user.tpl"
	
}
//post请求
func (c *UserController)DoAddUser(){
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println(username,password)
	c.Ctx.WriteString("用户中心")

}
//定义一个User的结构体i
/* type User struct{
	Username string `form:"username"`
	Password string `form:"password"`
	Hobby []string `form:"hobby"` 
}
func (c *UserController)DoEditUser(){
	u := User{}
	if err := c.ParseForm(&u);err != nil{//获取网页的form表单信息
		c.Ctx.WriteString("提交失败")
		return
	}
	fmt.Printf("%#v",u)
	c.Ctx.WriteString("解析Post数据成功")

}
func (c *UserController)GetUser(){
	u := User{
		Username: "张三",
		Password: "123456",
		Hobby: []string{"1,","2"},
	}
	//返回一个json数据
	c.Data["json"] = u
	c.ServeJSON()

} */