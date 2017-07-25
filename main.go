package main

import (
	_ "github.com/ss1917/my_cmdb/routers"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
