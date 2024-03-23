package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateUsersTable_20240323_111452 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsersTable_20240323_111452{}
	m.Created = "20240323_111452"

	migration.Register("CreateUsersTable_20240323_111452", m)
}

// Run the migrations
func (m *CreateUsersTable_20240323_111452) Up() {
	m.SQL("CREATE TABLE users (id bigint unsigned NOT NULL AUTO_INCREMENT, username varchar(64) DEFAULT NULL, password varchar(255) DEFAULT NULL, created_at datetime(3) DEFAULT NULL, updated_at datetime(3) DEFAULT NULL, PRIMARY KEY (id))")

}

// Reverse the migrations
func (m *CreateUsersTable_20240323_111452) Down() {
	m.SQL("DROP TABLE users")
}
