package main

import (
	"beego_todos/models"
	_ "beego_todos/routers"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web/context"

	"github.com/beego/beego/v2/client/orm"
	web "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var AuthFilter = func(ctx *context.Context) {
	userID := ctx.Input.Session("user_id")
	if userID == nil {
		if strings.HasPrefix(ctx.Input.URI(), "/todos") {
			ctx.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		}
	} else {
		ctx.Input.SetData("LoggedIn", true)
		user := models.UserFind(userID.(uint64))
		ctx.Input.SetData("username", user.Username)
	}
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlDataSource, _ := web.AppConfig.String("mysqlDataSource")
	orm.RegisterDataBase("default", "mysql", mysqlDataSource)
	orm.RegisterModel(new(models.Todo))
	orm.RegisterModel(new(models.User))

	web.InsertFilter("*", web.BeforeRouter, AuthFilter)
}

func main() {
	web.Run()
}
