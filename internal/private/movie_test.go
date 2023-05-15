package private

import (
	"database/sql"
	"os"
	"testing"
)

func TestCreateMovie(t *testing.T) {
	os.Remove("./data.db")
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		t.Fatalf("failt to open database, %v", err)
	}

	defer db.Close()

	m := &Movie{db: db}

	name := "test name"
	description := "test description"
	year := 1
	categoryID := "test categoryID"
	movie, err := m.Create(name, description, year, categoryID)

	if err != nil {
		t.Fatalf("fail to create Movie, %v", err)
	}

	if movie.Name != name {
		t.Fatalf("movies name is %s, expected %s", movie.Name, name)
	}

	if movie.Description != description {
		t.Fatalf("movies description is %s, expected %s", movie.Description, description)
	}

	if movie.Year != year {
		t.Fatalf("movies year is %v, expected %v", movie.Year, year)
	}

	if movie.CategoryID != categoryID {
		t.Fatalf("movies category_id is %s, expected %s", movie.CategoryID, categoryID)
	}
}
