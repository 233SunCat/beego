package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//
type Txt struct {
	Id       int
	Title    string
	Type     string
	Content  string    `orm:"size(10000)"`
	CreateAt time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8&loc=Local")

	orm.RegisterModel(new(Txt))
	//创建表格
	orm.RunSyncdb("default", false, true)

}

//insert
func (a *Txt) Insert(txt *Txt) {
	o := orm.NewOrm()
	id, err := o.Insert(txt)
	fmt.Print(id, err)
	if err != nil {

	}
}
func (a *Txt) Update(txt *Txt) error {
	if _, err := orm.NewOrm().Update(txt); err != nil {
		return err
	}
	return nil
}

//delete
func (a *Txt) Delete() error {
	if _, err := orm.NewOrm().Delete(a); err != nil {
		return err
	}
	return nil
}

//search
func (a *Txt) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(a)
}

//刷新数据库id
func (a *Txt) ReFresh() {
	orm1 := orm.NewOrm()
	_, err1 := orm1.Raw("ALTER TABLE `txt` DROP `id`;").Exec()
	_, err2 := orm1.Raw("ALTER TABLE `txt` ADD `id` int NOT NULL FIRST;").Exec()
	_, err3 := orm1.Raw("ALTER TABLE `txt` MODIFY COLUMN `id` int NOT NULL AUTO_INCREMENT,ADD PRIMARY KEY(id);").Exec()

	if err1 != nil || err2 != nil || err3 != nil {
		beego.Info("执行sql语句失败")
		return
	}
}
