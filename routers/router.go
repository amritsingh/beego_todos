package routers

import (
	"beego_todos/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/todos", &controllers.TodosController{}, "get:TodosIndex")
	beego.Router("/todos/new", &controllers.TodosController{}, "get:TodosNewForm")
	beego.Router("/todos", &controllers.TodosController{}, "post:TodosCreate")
	beego.Router("/todos/:id([0-9]+)", &controllers.TodosController{}, "get:TodosShow")
	beego.Router("/todos/edit/:id([0-9]+)", &controllers.TodosController{}, "get:TodosEditPage")
	beego.Router("/todos/:id", &controllers.TodosController{}, "post:TodosUpdate")
	beego.Router("/todos/:id/complete", &controllers.TodosController{}, "post:TodosComplete")
	beego.Router("/todos/:id", &controllers.TodosController{}, "delete:TodosDelete")

	beego.Router("/signup", &controllers.SessionsController{}, "get:SignupPage")
	beego.Router("/login", &controllers.SessionsController{}, "get:LoginPage")

	beego.Router("/signup", &controllers.SessionsController{}, "post:Signup")
	beego.Router("/login", &controllers.SessionsController{}, "post:Login")
	beego.Router("/logout", &controllers.SessionsController{}, "post:Logout")
}
