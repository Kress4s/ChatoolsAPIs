package routers

import (
	"ChatoolsAPIs/app/main/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/getoken", &controllers.UserController{}, "post:GenerateAuthorization")
	ns1 := beego.NewNamespace("/v1",
		beego.NSNamespace("/bot",
			beego.NSInclude(&controllers.BotController{}),
		),
		beego.NSNamespace("/contact",
			beego.NSInclude(&controllers.ContactController{}),
		),
	)
	beego.AddNamespace(ns1)
}
