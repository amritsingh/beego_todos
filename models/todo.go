package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Todo struct {
	Id          uint64
	Title       string
	Detail      string
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (n *Todo) TableName() string {
	return "todos"
}

func TodosGetAll() *[]*Todo {
	o := orm.NewOrm()
	var todos []*Todo
	// TODO:
	// Query the table todos (using model Todo) into "todos" variable
	// Query should fetch all records where deleted_at is NULL and sort by updated_at descending

	return &todos
}

func TodosCreate(title string, detail string) {
	o := orm.NewOrm()
	// TODO:
	// Insert a new record in the todos table.
	// Set the title and detail columns to the parameters passed to this function.
	// Don't forget to set created_at and updated_at timestamps.
}

func TodosFind(id uint64) *Todo {
	o := orm.NewOrm()
	todo := Todo{Id: id}
	err := o.Read(&todo)
	if err != nil {
		return nil
	} else {
		return &todo
	}
}

func (todo *Todo) Update(title string, detail string) {
	o := orm.NewOrm()
	// TODO:
	// Update the values of title, detail and updated_at
}

func TodosMarkComplete(id uint64, complete bool) {
	o := orm.NewOrm()
	// TODO:
	// If the "complete" variable is true, set completed_at column to current time
	// else set it to NULL. This will mark the TODO item completed or not completed.
	// Don't foget to set updated_at to current time.
}

func TodosMarkDelete(id uint64) {
	o := orm.NewOrm()
	// TODO:
	// To mark a TODO deleted, set deleted_at to current time.
	// Don't foget to set updated_at to current time.
	todo := TodosFind(id)
	todo.DeletedAt = time.Now()
	todo.UpdatedAt = time.Now()
	o.Update(todo)
}
