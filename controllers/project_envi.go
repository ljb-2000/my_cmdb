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
type ProjectEnviController struct {
	beego.Controller
}

// @router /:project [get]
func (this *ProjectEnviController) Get() {
	project := this.GetString(":project")

	if project != "" {
		if envi_info, err := models.Getprojectenvi(project); err == nil {
			this.Data["json"] = map[string]interface{}{
				"project":   project,
				"envi_list": envi_info,
				"status":    0,
				"msg":       "success",
			}
			this.ServeJSON()
		}
	}

	this.Data["json"] = map[string]interface{}{
		"project":   project,
		"envi_list": "",
		"status":    -1,
		"msg":       "获取项目失败",
	}
	this.ServeJSON()
}

// @router / [get]
func (this *ProjectEnviController) GetAll() {

	if proj_info, err := models.Getprojectproj(); err == nil {
		this.Data["json"] = map[string]interface{}{
			"proj_list": proj_info,
			"status":    0,
			"msg":       "success",
		}
		this.ServeJSON()
	}

	this.Data["json"] = map[string]interface{}{
		"envi_list": "",
		"status":    -1,
		"msg":       "获取项目失败",
	}
	this.ServeJSON()
}


// @router / [post]
func (this *ProjectEnviController) Post() {
	type ProjectInfo struct {
		Project string `json:"project"`
		Envi    string `json:"envi"`
	}

	var (
		projectinfo *ProjectInfo = new(ProjectInfo)
		//conf                     = new(models.ProjectEnvi)
	)
	json.Unmarshal(this.Ctx.Input.RequestBody, projectinfo)
	fmt.Println(projectinfo)

	newconf := &models.ProjectEnvi{
		Project: projectinfo.Project,
		Envi:    projectinfo.Envi,
		Uptime:  time.Now(),
	}
	if created, err := models.InsertProjectenvi(newconf); err == nil && created {

		this.Data["json"] = map[string]interface{}{
			"status":  0,
			"msg":     "添加环境成功",
		}
		this.ServeJSON()

	} else {
		this.Data["json"] = map[string]interface{}{
			"status": -1,
			"msg":    "添加环境失败",
		}
		this.ServeJSON()
	}
}

// @router /:pid [delete]
func (this *ProjectEnviController) Delete() {
	pid := this.GetString(":pid")


	if Id, erro := strconv.Atoi(pid); erro != nil {
		this.Data["json"] = map[string]interface{}{
			"status": -1,
			"msg":    "请输入正确的ID",
		}
		this.ServeJSON()

	} else {
		if err := models.DeleteProjectEnvi(Id); err == nil {
			this.Data["json"] = map[string]interface{}{
				"status": 0,
				"msg":    "删除环境成功",
			}
			this.ServeJSON()
		}
	}

	this.Data["json"] = map[string]interface{}{
		"status": -2,
		"msg":    "删除环境失败",
	}
	this.ServeJSON()
}