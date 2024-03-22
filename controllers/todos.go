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
	// Get all todos
	todos := models.TodosGetAll()
	c.Data["todos"] = todos
	c.TplName = "todos/index.tpl"
}

func (c *TodosController) TodosNewForm() {
	c.TplName = "todos/new.tpl"
}

func (c *TodosController) TodosCreate() {
	// TODO:
	// Read the form parametes "title" and "detail".
	// Call models.TodosCreate function to create a new TODO item

	c.Redirect("/todos", http.StatusFound)
}

func (c *TodosController) TodosShow() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	// TODO:
	// Retrieve the TODO item from the identifier "id"
	// Set the retrieved todo item to c.Data["todo"]
	c.TplName = "todos/show.tpl"
}

func (c *TodosController) TodosEditPage() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	// TODO:
	// Retrieve the TODO item from the identifier "id"
	// Set the retrieved todo item to c.Data["todo"]
	c.TplName = "todos/edit.tpl"
}

func (c *TodosController) TodosUpdate() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	todo := models.TodosFind(id)
	title := c.GetString("title")
	detail := c.GetString("detail")
	// TODO:
	// Call Update() function defined in the model to update the record.
	// Redirect to the show page - /todos/(id)
}

func (c *TodosController) TodosDelete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	models.TodosMarkDelete(id)
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
}

func (c *TodosController) TodosComplete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	state := c.Ctx.Input.Param(":state")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	models.TodosMarkComplete(id, state == "true")
	c.Ctx.ResponseWriter.WriteHeader(http.StatusAccepted)
}
