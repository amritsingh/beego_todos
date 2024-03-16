package models

import (
	"fmt"
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
	numRows, err := o.QueryTable(new(Todo)).
		Filter("deleted_at__isnull", true).
		OrderBy("-updated_at").
		All(&todos)
	if err != nil {
		panic(err)
	}
	fmt.Println(numRows)
	return &todos
}

func TodosCreate(title string, detail string) {
	o := orm.NewOrm()
	currTime := time.Now()
	todo := Todo{Title: title, Detail: detail, CreatedAt: currTime, UpdatedAt: currTime}
	id, err := o.Insert(&todo)
	fmt.Println(id)
	fmt.Println(err)
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
	todo.Title = title
	todo.Detail = detail
	todo.UpdatedAt = time.Now()
	o.Update(todo)
}

func TodosMarkComplete(id uint64, complete bool) {
	o := orm.NewOrm()
	todo := TodosFind(id)
	if complete {
		*todo.CompletedAt = time.Now()
	} else {
		todo.CompletedAt = nil
	}
	todo.UpdatedAt = time.Now()
	o.Update(todo)
}

func TodosMarkDelete(id uint64) {
	o := orm.NewOrm()
	todo := TodosFind(id)
	todo.DeletedAt = time.Now()
	todo.UpdatedAt = time.Now()
	o.Update(todo)
}
