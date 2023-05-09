package private

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateCategory(t *testing.T) {
	os.Remove("./data.db")
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		t.Fatalf("failted to open database, %v", err)
	}
	defer db.Close()

	c := &Category{db: db}

	name := "test category"
	description := "test category description"
	category, err := c.Create(name, description)

	if err != nil {
		t.Fatalf("failed to create category, %v", err)
	}

	if category.Name != name {
		t.Errorf("category name is %s; expected %s", category.Name, name)
	}

	if category.Description != description {
		t.Errorf("category description is %s; expected %s", category.Description, description)
	}
}
