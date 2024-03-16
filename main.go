package main

import (
	"beego_todos/models"
	_ "beego_todos/routers"

	"github.com/beego/beego/v2/client/orm"
	web "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlDataSource, _ := web.AppConfig.String("mysqlDataSource")
	orm.RegisterDataBase("default", "mysql", mysqlDataSource)
	orm.RegisterModel(new(models.Todo))
}

func main() {
	web.Run()
}
