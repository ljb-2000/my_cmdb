// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/ss1917/my_cmdb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/project_conf",
			beego.NSNamespace("/info",
				beego.NSInclude(
					&controllers.ProjectConfController{},
				),

			),
			beego.NSNamespace("/envi",
				beego.NSInclude(
					&controllers.ProjectEnviController{},
				),
			),
			beego.NSNamespace("/etcd",
				beego.NSInclude(
					&controllers.ProjectSetEtcdController{},
				),
			),
		),
	)
	beego.AddNamespace(ns)
}
