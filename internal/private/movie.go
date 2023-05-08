package private

import (
	"database/sql"

	"github.com/google/uuid"
)

type Movie struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	Year        int
	CategoryID  string
}

func NewMovie(db *sql.DB) *Movie {
	return &Movie{db: db}
}

func (m *Movie) Create(name string, description string, year int, categoryID string) (*Movie, error) {
	id := uuid.New().String()
	_, err := m.db.Exec("INSERT INTO movies (id, name, description, year, category_id) VALUES ($1, $2, $3, $4, $5)",
		id, name, description, year, categoryID)
	if err != nil {
		return nil, err
	}
	return &Movie{
		ID:          id,
		Name:        name,
		Description: description,
		Year:        year,
		CategoryID:  categoryID,
	}, nil
}

func (m *Movie) FindAll() ([]Movie, error) {
	rows, err := m.db.Query("SELECT id, name, description, year, category_id FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	movies := []Movie{}
	for rows.Next() {
		var id string
		var name string
		var description string
		var year int
		var categoryID string

		if err := rows.Scan(&id, &name, &description, &year, &categoryID); err != nil {
			return nil, err
		}
		movies = append(movies, Movie{ID: id, Name: name, Description: description, Year: year, CategoryID: categoryID})
	}
	return movies, nil
}

func (m *Movie) FindByCategoryID(categoryID string) ([]Movie, error) {
	rows, err := m.db.Query("SELECT id, name, description, year, category_id FROM movies WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	movies := []Movie{}
	for rows.Next() {
		var id string
		var name string
		var description string
		var year int
		var categoryID string

		if err = rows.Scan(&id, &name, &description, &year, &categoryID); err != nil {
			return nil, err
		}
		movies = append(movies, Movie{ID: id, Name: name, Description: description, Year: year, CategoryID: categoryID})
	}
	return movies, nil
}
