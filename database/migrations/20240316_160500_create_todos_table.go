package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateTodosTable_20240316_160500 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTodosTable_20240316_160500{}
	m.Created = "20240316_160500"

	migration.Register("CreateTodosTable_20240316_160500", m)
}

// Run the migrations
func (m *CreateTodosTable_20240316_160500) Up() {
	m.SQL("CREATE TABLE todos (id bigint unsigned NOT NULL AUTO_INCREMENT, title varchar(255) DEFAULT NULL, detail text, completed_at datetime(3) DEFAULT NULL, created_at datetime(3) DEFAULT NULL, updated_at datetime(3) DEFAULT NULL, deleted_at datetime(3) DEFAULT NULL, PRIMARY KEY (id), KEY idx_todos_updated_at (updated_at), KEY idx_todos_deleted_at (deleted_at))")
}

// Reverse the migrations
func (m *CreateTodosTable_20240316_160500) Down() {
	m.SQL("DROP TABLE todos")
}
