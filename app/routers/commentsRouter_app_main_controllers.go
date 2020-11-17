package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:BotController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:BotController"],
        beego.ControllerComments{
            Method: "GetQR",
            Router: "/getqrcode",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:ContactController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:ContactController"],
        beego.ControllerComments{
            Method: "Betch",
            Router: "/betch",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:ContactController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:ContactController"],
        beego.ControllerComments{
            Method: "ListAll",
            Router: "/list/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:ContactController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:ContactController"],
        beego.ControllerComments{
            Method: "ListGroup",
            Router: "/list/group",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:ContactController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:ContactController"],
        beego.ControllerComments{
            Method: "Search",
            Router: "/search",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
