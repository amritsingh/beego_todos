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
	UserId      uint64
}

func (n *Todo) TableName() string {
	return "todos"
}

func TodosGetAll(user *User) *[]*Todo {
	o := orm.NewOrm()
	var todos []*Todo
	numRows, err := o.QueryTable(new(Todo)).
		Filter("user_id", user.Id).
		Filter("deleted_at__isnull", true).
		OrderBy("-updated_at").
		All(&todos)
	if err != nil {
		panic(err)
	}
	fmt.Println(numRows)
	return &todos
}

func TodosCreate(user *User, title string, detail string) {
	o := orm.NewOrm()
	currTime := time.Now()
	todo := Todo{Title: title, Detail: detail, CreatedAt: currTime, UpdatedAt: currTime, UserId: user.Id}
	id, err := o.Insert(&todo)
	fmt.Println(id)
	fmt.Println(err)
}

func TodosFind(user *User, id uint64) *Todo {
	o := orm.NewOrm()
	todo := Todo{Id: id, UserId: user.Id}
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

func TodosMarkComplete(user *User, id uint64, complete bool) {
	o := orm.NewOrm()
	todo := TodosFind(user, id)
	if complete {
		*todo.CompletedAt = time.Now()
	} else {
		todo.CompletedAt = nil
	}
	todo.UpdatedAt = time.Now()
	o.Update(todo)
}

func TodosMarkDelete(user *User, id uint64) {
	o := orm.NewOrm()
	todo := TodosFind(user, id)
	todo.DeletedAt = time.Now()
	todo.UpdatedAt = time.Now()
	o.Update(todo)
}
