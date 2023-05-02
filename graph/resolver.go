package graph

import "github.com/lanpaiva/movies-graphql/internal/private"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *private.Category
}
