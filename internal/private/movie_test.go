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

func TestFindAllMovies(t *testing.T) {
	os.Remove("./data.db")
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		t.Fatalf("Fail to conoct at database %v", err)
	}
	defer db.Close()

	m := &Movie{db: db}

	_, err = m.Create("movie1", "description1", 1, "id1")
	if err != nil {
		t.Fatalf("fail to create movie1 %v", err)
	}

	_, err = m.Create("movie2", "description2", 2, "id2")
	if err != nil {
		t.Fatalf("failt to create movie2 %v", err)
	}

	movies, err := m.FindAll()

	if err != nil {
		t.Fatalf("fail to find movies %v", err)
	}

	if len(movies) != 2 {
		t.Errorf("got %d movies, expected 2", len(movies))
	}

	if movies[0].Name != "movie1" || movies[0].Description != "description1" {
		t.Errorf("name or description is different, got %v; expected {Name: 'movie1', Description: 'description1'}", movies[0])
	}

	if movies[1].Name != "movie2" || movies[1].Description != "description2" {
		t.Errorf("name or description is different, got %v; expected {Name: 'movie2', Description: 'description2'}", movies[1])
	}
}
