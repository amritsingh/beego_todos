package controllers

import (
	"beego_todos/models"
	"fmt"
	"net/http"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

// TodosController operations for Todos
type TodosController struct {
	beego.Controller
}

func (c *TodosController) TodosIndex() {
	user := GetUserFromSession(&c.Controller)
	// Get all todos
	todos := models.TodosGetAll(user)
	c.Data["todos"] = todos
	c.TplName = "todos/index.tpl"
}

func (c *TodosController) TodosNewForm() {
	c.TplName = "todos/new.tpl"
}

func (c *TodosController) TodosCreate() {
	user := GetUserFromSession(&c.Controller)
	title := c.GetString("title")
	detail := c.GetString("detail")

	models.TodosCreate(user, title, detail)
	c.Redirect("/todos", http.StatusFound)
}

func (c *TodosController) TodosShow() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	user := GetUserFromSession(&c.Controller)
	todo := models.TodosFind(user, id)
	c.Data["todo"] = todo
	c.TplName = "todos/show.tpl"
}

func (c *TodosController) TodosEditPage() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	user := GetUserFromSession(&c.Controller)
	todo := models.TodosFind(user, id)
	c.Data["todo"] = todo
	c.TplName = "todos/edit.tpl"
}

func (c *TodosController) TodosUpdate() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	user := GetUserFromSession(&c.Controller)
	todo := models.TodosFind(user, id)
	title := c.GetString("title")
	detail := c.GetString("detail")
	todo.Update(title, detail)
	c.Redirect("/todos/"+idStr, http.StatusFound)
}

func (c *TodosController) TodosDelete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	user := GetUserFromSession(&c.Controller)
	models.TodosMarkDelete(user, id)
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
}

func (c *TodosController) TodosComplete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	state := c.Ctx.Input.Param(":state")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	user := GetUserFromSession(&c.Controller)
	models.TodosMarkComplete(user, id, state == "true")
	c.Ctx.ResponseWriter.WriteHeader(http.StatusAccepted)
}
