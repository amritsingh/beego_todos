package controllers

import (
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	user := GetUserFromSession(&c.Controller)
	if user.Id == 0 {
		c.Data["Title"] = "TODO List App"
		c.TplName = "index.tpl"
	} else {
		c.Redirect("/todos/", http.StatusFound)
	}
}
