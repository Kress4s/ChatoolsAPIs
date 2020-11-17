package dbmysql

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//导入数据库引擎

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqladdr := beego.AppConfig.String("MYSQL::IP")
	mysqlport := beego.AppConfig.String("MYSQL::PORT")
	mysqluser := beego.AppConfig.String("MYSQL::USER")
	mysqlpasswd := beego.AppConfig.String("MYSQL::PWD")
	masterdb := beego.AppConfig.String("MYSQL::DB_NAME")
	mysqlcharset := beego.AppConfig.String("MYSQL::CHARSET")
	loc := beego.AppConfig.String("MYSQL::LOC")
	connectStr := mysqluser + ":" + mysqlpasswd + "@" + "tcp" + "(" + mysqladdr + ":" + mysqlport + ")/" + masterdb +
		"?charset=" + mysqlcharset + "&loc=" + loc
	orm.RegisterDataBase("default", "mysql", connectStr)
}
