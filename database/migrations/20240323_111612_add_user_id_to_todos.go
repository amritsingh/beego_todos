package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddUserIdToTodos_20240323_111612 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddUserIdToTodos_20240323_111612{}
	m.Created = "20240323_111612"

	migration.Register("AddUserIdToTodos_20240323_111612", m)
}

// Run the migrations
func (m *AddUserIdToTodos_20240323_111612) Up() {
	m.SQL("ALTER TABLE todos ADD user_id bigint")
}

// Reverse the migrations
func (m *AddUserIdToTodos_20240323_111612) Down() {
	m.SQL("ALTER TABLE todos DROP COLUMN user_id")
}
