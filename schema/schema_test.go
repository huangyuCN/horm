package schema

import (
	"horm/dialect"
	"testing"
)

type User struct {
	Name string `horm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse user struct")
	}
	if schema.GetFiled("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
