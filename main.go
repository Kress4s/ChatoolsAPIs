package main

import (
	"ChatoolsAPIs/app/bridage/exception"
	_ "ChatoolsAPIs/app/bridage/models"
	"ChatoolsAPIs/app/bridage/path"
	_ "ChatoolsAPIs/app/common/dbmysql"
	_ "ChatoolsAPIs/app/routers"
	"fmt"
	"os"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	var checkAcess = func(ctx *context.Context) {
		if token := ctx.Input.Header("Authorization"); token == "" {
			ctx.Input.RunController = reflect.TypeOf(exception.GetInst())
			ctx.Input.RunMethod = "ExceptToken"
		}
	}
	//允许跨站访问
	//if beego.BConfig.RunMode == "dev" {
	beego.InsertFilter("/v1/*", beego.BeforeRouter, checkAcess)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Token", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//}
	createDIR() //初始化必要目录
	//设置日志规则
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/chatools.log","separate":["error", "warning", "notice", "info", "debug"]}`)
	logs.EnableFuncCallDepth(true)
	beego.AddViewPath("template")
	args := os.Args //获取用户输入的所有参数
	if args == nil || len(args) < 2 {
		//如果用户没有输入,或参数个数不够,则调用该函数提示用户
		//beego.Run()
		cmdUsage()
	} else if len(args) == 2 {
		switch args[1] {
		case "start":
			// if common.DetectGRPC() == constant.GRPC_RUN_SERVER {
			beego.Run()
			// }
		case "orm":
			orm.RunCommand()
		default:
			cmdUsage()
		}
	} else if len(args) > 2 {
		if args[1] == "orm" {
			orm.RunCommand()
		} else {
			cmdUsage()
		}
	}
}

//cmdUsage 显示命令行帮助
func cmdUsage() {
	fmt.Println(`
USAGE
	betaSrv [commond]
AVAILABLE COMMANDS
	start                     Start beta managerment server node.
	orm                       Operate the database.
	`)
}

//初始化公共目录
func createDIR() {
	var err error
	//初始化日志目录
	if _, err = os.Stat(path.GetLogsPath()); err != nil {
		os.MkdirAll(path.GetLogsPath(), os.ModePerm)
	}

	if _, err = os.Stat(path.GetUploadPath()); err != nil {
		os.MkdirAll(path.GetUploadPath(), os.ModePerm)
	}
}
