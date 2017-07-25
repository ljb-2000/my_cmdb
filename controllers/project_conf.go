package controllers

import (
	"github.com/ss1917/my_cmdb/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"time"
	"fmt"
	"strconv"
)

// Operations about object
type ProjectConfController struct {
	beego.Controller
}

// @Title Create
// @Description create project
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *ProjectConfController) Post() {
	type ProjectInfo struct {
		Project string `json:"project"`
		Envi    string `json:"envi"`
		Key     string `json:"key"`
		Value   string `json:"value"`
		Notes   string `json:"notes"`
	}

	var (
		projectinfo *ProjectInfo = new(ProjectInfo)
		conf                     = new(models.ProjectConf)
	)
	json.Unmarshal(this.Ctx.Input.RequestBody, projectinfo)
	fmt.Println(projectinfo)

	newconf := &models.ProjectConf{
		Project: projectinfo.Project,
		Envi:    projectinfo.Envi,
		Key:     projectinfo.Key,
		Value:   projectinfo.Value,
		Notes:   projectinfo.Notes,
		Uptime:  time.Now(),
	}
	if created, err := conf.InsertProject(newconf); err == nil && created {
		this.Data["json"] = map[string]interface{}{
			"status":  0,
			"msg":     "success",
		}
		this.ServeJSON()

	} else {
		this.Data["json"] = map[string]interface{}{
			"status": -1,
			"msg":    "添加失败",
		}
		this.ServeJSON()
	}
}

// @Title Get
// @Description find project by project environment
// @Param	project		path 	string	true		"the objectid you want to get"
// @Success 200 {project} models.Object
// @Failure 403 :project is empty
// @router /:project/:envi [get]
func (this *ProjectConfController) Get() {
	project := this.Ctx.Input.Param(":project")
	envi := this.Ctx.Input.Param(":envi")
	type ProInfo struct {
		Key    string    `json:"key"`
		Value  string    `json:"value"`
		Notes  string    `json:"notes"`
		Uptime time.Time `json:"uptime"`
	}
	var (
		conf = new(models.ProjectConf)
	)
	if project == "" || envi == "" {
		this.Data["json"] = map[string]interface{}{
			"project": project,
			"envi":    envi,
			"status":  -1,
			"msg":     "项目和环境名不能为空",
		}
		this.ServeJSON()
	}

	if proj_info, err := conf.Getproject(project, envi); err == nil {
		this.Data["json"] = map[string]interface{}{
			"project":   project,
			"envi":      envi,
			"status":    0,
			"msg":       "success",
			"proj_info": proj_info,
		}
		this.ServeJSON()

	} else {
		this.Data["json"] = map[string]interface{}{
			"project": project,
			"envi":    envi,
			"status":  -2,
			"msg":     "获取配置失败",
		}
		this.ServeJSON()

	}
}

// @router /:pid [delete]
func (this *ProjectConfController) Delete() {
	pid := this.GetString(":pid")
	var (
		conf = new(models.ProjectConf)
	)

	if Id, erro := strconv.Atoi(pid); erro != nil {
		this.Data["json"] = map[string]interface{}{
			"status": -1,
			"msg":    "请输入正确的ID",
		}
		this.ServeJSON()

	} else {
		if err := conf.DeleteProject(Id); err == nil {
			this.Data["json"] = map[string]interface{}{
				"status": 0,
				"msg":    "删除成功",
			}
			this.ServeJSON()
		}
	}

	this.Data["json"] = map[string]interface{}{
		"status": -2,
		"msg":    "删除失败",
	}
	this.ServeJSON()
}

// @router / [patch]
func (this *ProjectConfController) Patch() {
	type ProjectInfo struct {
		Pid   string `json:"pid"`
		Value string `json:"value"`
		Notes string `json:"notes"`
	}

	var (
		projectinfo *ProjectInfo = new(ProjectInfo)
		conf                     = new(models.ProjectConf)
	)

	json.Unmarshal(this.Ctx.Input.RequestBody, projectinfo)
	fmt.Println(projectinfo)

	if Id, erro := strconv.Atoi(projectinfo.Pid); erro != nil {
		this.Data["json"] = map[string]interface{}{
			"status": -1,
			"msg":    "请输入正确的ID",
		}
		this.ServeJSON()

	} else {

		if _, err := conf.PatchProject(Id, projectinfo.Value, projectinfo.Notes); err == nil {
			this.Data["json"] = map[string]interface{}{
				"status": 0,
				"msg":    "更新成功",
			}
			this.ServeJSON()

		} else {
			this.Data["json"] = map[string]interface{}{
				"status": -2,
				"msg":    "更新失败",
			}
			this.ServeJSON()

		}
	}
}
