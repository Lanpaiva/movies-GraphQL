package model

type Movie struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Year        *int    `json:"year,omitempty"`
}
