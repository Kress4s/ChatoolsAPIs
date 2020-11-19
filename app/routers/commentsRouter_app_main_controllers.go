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

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:BotController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:BotController"],
        beego.ControllerComments{
            Method: "GetMyInfo",
            Router: "/profile",
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

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/create",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetGroupInfoA",
            Router: "/get/detail",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetGroupInfoB",
            Router: "/get/info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GetGroupMembers",
            Router: "/get/members",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"],
        beego.ControllerComments{
            Method: "AddGroupMember",
            Router: "/members/add",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"],
        beego.ControllerComments{
            Method: "DeleteGroupMember",
            Router: "/members/delete",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"],
        beego.ControllerComments{
            Method: "QuitGroup",
            Router: "/quit",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:GroupController"],
        beego.ControllerComments{
            Method: "SetAnnounceMent",
            Router: "/set/announcement",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:MessageController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:MessageController"],
        beego.ControllerComments{
            Method: "SendTheImage",
            Router: "/send/image",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:MessageController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:MessageController"],
        beego.ControllerComments{
            Method: "SendTheText",
            Router: "/send/text",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:MessageController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:MessageController"],
        beego.ControllerComments{
            Method: "SendTheVideo",
            Router: "/send/video",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:MessageController"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:MessageController"],
        beego.ControllerComments{
            Method: "SyncRecieveMessage",
            Router: "/sync/message/pushstream",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:SnsControllers"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:SnsControllers"],
        beego.ControllerComments{
            Method: "ListMyFriendSns",
            Router: "/list/friend",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:SnsControllers"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:SnsControllers"],
        beego.ControllerComments{
            Method: "ListMe",
            Router: "/list/me",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:SnsControllers"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:SnsControllers"],
        beego.ControllerComments{
            Method: "SendImageAndTextSns",
            Router: "/send/image",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:SnsControllers"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:SnsControllers"],
        beego.ControllerComments{
            Method: "SendTheTextSns",
            Router: "/send/text",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:SnsControllers"] = append(beego.GlobalControllerRouter["ChatoolsAPIs/app/main/controllers:SnsControllers"],
        beego.ControllerComments{
            Method: "SendTheVideoSns",
            Router: "/send/video",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
