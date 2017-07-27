package controllers

import (
	"github.com/ss1917/my_cmdb/models"
	"bytes"
	"github.com/astaxie/beego"
	"encoding/json"
	"time"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"github.com/astaxie/beego/logs"
)

// Operations about object
type ProjectSetEtcdController struct {
	beego.Controller
}

// @Title Submit configuration
// @Description Submit configuration
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {json} {"status": 0,"msg": "success"}
// @Failure 403 body is empty
// @router / [post]
func (this *ProjectSetEtcdController) Post() {
	type ProjectInfo struct {
		Project string `json:"project"`
		Envi    string `json:"envi"`
	}

	//beego.AppConfig.String("etcd" + "::address")
	var (
		projectinfo *ProjectInfo = new(ProjectInfo)
		conf                     = new(models.ProjectConf)
	)
	json.Unmarshal(this.Ctx.Input.RequestBody, projectinfo)

	cfg := client.Config{
		Endpoints: []string{beego.AppConfig.String("etcd::address01"),
							beego.AppConfig.String("etcd::address02")},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		logs.Error("etcd client err",err)
	}
	kapi := client.NewKeysAPI(c)

	if proj_info, err := conf.Getproject(projectinfo.Project, projectinfo.Envi); err == nil {
		for _, c := range proj_info {
			//s := fmt.Sprintf("%s~~%s~~%s", c["Project"], c["Envi"], c["Key"])
			b := bytes.Buffer{}
			b.WriteString("/my_project/")
			b.WriteString(c["Project"].(string))
			b.WriteString("/")
			b.WriteString(c["Envi"].(string))
			b.WriteString("/")
			b.WriteString(c["Key"].(string))
			s := b.String()
			if _, err := kapi.Set(context.Background(), s, c["Value"].(string), nil); err != nil {
				logs.Error("set etcd", err)
				this.Data["json"] = map[string]interface{}{
					"status": -1,
					"msg":    "写入ETCD失败",
				}
				this.ServeJSON()
			} else {
				logs.Debug("Set is done. Key is %q\n", s)
			}
		}

		this.Data["json"] = map[string]interface{}{
			"status": 0,
			"msg":    "success",
		}
		this.ServeJSON()

	} else {
		this.Data["json"] = map[string]interface{}{
			"status": -2,
			"msg":    "添加失败",
		}
		this.ServeJSON()
	}
}
