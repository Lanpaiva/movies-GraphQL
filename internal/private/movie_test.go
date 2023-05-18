package private

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateMovie(t *testing.T) {
	os.Remove("./data.db")
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		t.Fatalf("fail to open database, %v", err)
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

func TestFindByCategoryID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error to create mock database: %v", err)
	}
	defer db.Close()

	movie := &Movie{db: db}

	categoryID := "123"
	expectedMovies := []Movie{
		{ID: "1", Name: "Movie 1", Description: "Description 1", Year: 2021, CategoryID: categoryID},
		{ID: "2", Name: "Movie 2", Description: "Description 2", Year: 2022, CategoryID: categoryID},
	}
	mock.ExpectQuery("SELECT id, name, description, year, category_id FROM movies").
		WithArgs(categoryID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "year", "category_id"}).
				AddRow(expectedMovies[0].ID, expectedMovies[0].Name, expectedMovies[0].Description, expectedMovies[0].Year, expectedMovies[0].CategoryID).
				AddRow(expectedMovies[1].ID, expectedMovies[1].Name, expectedMovies[1].Description, expectedMovies[1].Year, expectedMovies[1].CategoryID),
		)

	movies, err := movie.FindByCategoryID(categoryID)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if len(movies) != len(expectedMovies) {
		t.Errorf("Number of movies is diferent. Expected: %d, Returned: %d", len(expectedMovies), len(movies))
	}

	for i, expected := range expectedMovies {
		if movies[i].ID != expected.ID {
			t.Errorf("ID of movie #%d is diferent. Expected: %s, Returned: %s", i, expected.ID, movies[i].ID)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectativas não atendidas: %v", err)
	}
}
