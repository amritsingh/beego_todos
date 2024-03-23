package main

import (
	"beego_todos/models"
	_ "beego_todos/routers"

	"github.com/beego/beego/v2/server/web/context"

	"github.com/beego/beego/v2/client/orm"
	web "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var AuthFilter = func(ctx *context.Context) {
	// TODO:
	// Step 1: Get the user ID from the session by accessing.
	// Step 2: If the user is not logged in:
	//         Check if the request URI starts with "/todos"
	// 		   If the request URI starts with "/todos",
	//         redirect the user to the login page ("/login").
	// Step 3: If the user is logged in:
	//         - Set the "LoggedIn" flag in the request context data to true.
	//         - Retrieve the user details from the database using the user ID by calling the appropriate model function.
	//         - Set the username in the request context data.
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlDataSource, _ := web.AppConfig.String("mysqlDataSource")
	orm.RegisterDataBase("default", "mysql", mysqlDataSource)
	orm.RegisterModel(new(models.Todo))
	orm.RegisterModel(new(models.User))

	// TODO:
	// Uncomment the following code to add the middlewar
	// web.InsertFilter("*", web.BeforeRouter, AuthFilter)
}

func main() {
	web.Run()
}
