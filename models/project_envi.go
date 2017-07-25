package models

import (
	"fmt"
	//"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego/orm"
)

type ProjectEnvi struct {
	Id      int       `orm:"column(pid);auto"`
	Project string    `orm:"column(project);size(80)"`
	Envi    string    `orm:"column(envi);size(20)"`
	Uptime  time.Time `orm:"column(uptime);type(datetime)"`
}

//定义数据表名
func (t *ProjectEnvi) TableName() string {
	return "project_envi"
}

func Getprojectenvi(Pro string) (maps []orm.Params, err error) {
	o := orm.NewOrm()
	//读操作用只读库
	o.Using("read")
	c := new(ProjectEnvi)
	fmt.Println(Pro)
	//var maps []orm.Params
	if _, err = o.QueryTable(c).Filter("Project", Pro).Values(&maps); err != nil {
		return
	}
	return
}

func Getprojectproj() (maps []orm.Params, err error) {
	o := orm.NewOrm()
	//读操作用只读库
	o.Using("read")
	c := new(ProjectEnvi)
	//var maps []orm.Params
	if _, err = o.QueryTable(c).GroupBy("Project").Values(&maps); err != nil {
		return
	}
	return
}

//添加
func InsertProjectenvi(newconf *ProjectEnvi) (created bool, err error) {
	o := orm.NewOrm()
	//开启事务
	o.Begin()
	if created, _, err = o.ReadOrCreate(newconf, "Project","Envi"); err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}

//通过ID删除
func DeleteProjectEnvi(Id int) (err error) {
	c := new(ProjectEnvi)
	_, err = orm.NewOrm().QueryTable(c).Filter("Id", Id).Delete()
	return
}