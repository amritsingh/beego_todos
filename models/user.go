package models

import (
	"beego_todos/helpers"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"size:64"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}

func UserCheckAvailability(email string) bool {
	o := orm.NewOrm()
	var user User
	o.QueryTable(new(User)).Filter("username", email).One(&user)
	return (user.Id == 0) // if ID == 0, email is not signed up, hence available
}

func UserCreate(email string, password string) *User {
	o := orm.NewOrm()
	hshPasswd, _ := helpers.HashPassword(password)
	entry := User{Username: email, Password: hshPasswd}
	o.Insert(&entry)
	return &entry
}

func UserFind(id uint64) *User {
	var user User
	o := orm.NewOrm()
	o.QueryTable(new(User)).Filter("id", id).One(&user)
	return &user
}

func UserCheck(email string, password string) *User {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable(new(User)).Filter("username", email).One(&user)
	fmt.Println(err)
	if user.Id == 0 {
		return nil
	}

	match := helpers.CheckPasswordHash(password, user.Password)
	fmt.Println(match)
	if match {
		return &user
	} else {
		return nil
	}
}
