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

func TestFindAll(t *testing.T) {
	os.Remove("./data.db")
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		t.Fatalf("failed to open database, %v", err)
	}
	defer db.Close()

	c := &Category{db: db}

	_, err = c.Create("category1", "description1")
	if err != nil {
		t.Fatalf("failed to create category, %v", err)
	}

	_, err = c.Create("category2", "description2")
	if err != nil {
		t.Fatalf("failed to create category, %v", err)
	}

	categories, err := c.FindAll()

	if err != nil {
		t.Fatalf("failed to find categories, %v", err)
	}

	if len(categories) != 2 {
		t.Errorf("got %d categories; expected 2", len(categories))
	}

	if categories[0].Name != "category1" || categories[0].Description != "description1" {
		t.Errorf("category 1 is incorrect; got %v, expected {Name: 'category1', Description: 'description1'}", categories[0])
	}

	if categories[1].Name != "category2" || categories[1].Description != "description2" {
		t.Errorf("category 2 is incorrect; got %v, expected {Name: 'category2', Description: 'description2'}", categories[1])
	}
}
