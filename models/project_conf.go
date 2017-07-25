package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	ProjectConfs map[string]*ProjectConf
)

type ProjectConf struct {
	Id      int       `orm:"column(cid);auto"`
	Project string    `orm:"column(project);size(80)"`
	Envi    string    `orm:"column(envi);size(20)"`
	Key     string    `orm:"column(key);size(50)"`
	Value   string    `orm:"column(value);size(800)"`
	Notes   string    `orm:"column(notes);size(300)"`
	Uptime  time.Time `orm:"column(uptime);type(datetime)"`
}

//定义数据表名
func (t *ProjectConf) TableName() string {
	return "project_conf"
}

func (this *ProjectConf) Getproject(Pro, Envi string) (maps []orm.Params, err error) {
	o := orm.NewOrm()
	//读操作用只读库
	o.Using("read")
	c := new(ProjectConf)
	//var maps []orm.Params
	if _, err = o.QueryTable(c).Filter("project", Pro).Filter("envi", Envi).Values(&maps); err != nil {
		return
	}
	return
}

//添加
func (this *ProjectConf) InsertProject(newconf *ProjectConf) (created bool, err error) {
	o := orm.NewOrm()
	//开启事务
	o.Begin()
	if created, _, err = o.ReadOrCreate(newconf, "Project", "Envi", "Key"); err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}

//通过ID删除
func (this *ProjectConf) DeleteProject(Id int) (err error) {
	c := new(ProjectConf)
	_, err = orm.NewOrm().QueryTable(c).Filter("Id", Id).Delete()
	return
}

//通过自增ID更新value
func (this *ProjectConf) PatchProject(Id int, value, notes string) (num int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	c := new(ProjectConf)
	if num, err = o.QueryTable(c).Filter("Id", Id).Update(orm.Params{"Value": value,
		"Notes":                                                              notes, "Uptime": time.Now()}); err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}
