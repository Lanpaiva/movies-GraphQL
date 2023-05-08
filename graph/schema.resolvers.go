package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"

	"github.com/lanpaiva/movies-graphql/graph/model"
)

// Movies is the resolver for the movies field.
func (r *categoryResolver) Movies(ctx context.Context, obj *model.Category) ([]*model.Movie, error) {
	movies, err := r.MovieDB.FindByCategoryID(obj.ID)
	if err != nil {
		return nil, err
	}
	var moviesModel []*model.Movie
	for _, movie := range movies {
		moviesModel = append(moviesModel, &model.Movie{
			ID:          movie.ID,
			Name:        movie.Name,
			Description: &movie.Description,
			Year:        &movie.Year,
		})
	}
	return moviesModel, nil
}

// Category is the resolver for the category field.
func (r *movieResolver) Category(ctx context.Context, obj *model.Movie) (*model.Category, error) {
	category, err := r.CategoryDB.FindByMovieID(obj.ID)
	if err != nil {
		return nil, err
	}
	return &model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: &category.Description,
	}, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category, err := r.CategoryDB.Create(input.Name, *input.Description)
	if err != nil {
		return nil, err
	}
	return &model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: &category.Description,
	}, nil
}

// CreateMovie is the resolver for the createMovie field.
func (r *mutationResolver) CreateMovie(ctx context.Context, input model.NewMovie) (*model.Movie, error) {
	movie, err := r.MovieDB.Create(input.Name, *input.Description, *input.Year, input.CategoryID)
	if err != nil {
		return nil, err
	}
	return &model.Movie{
		ID:          movie.ID,
		Name:        movie.Name,
		Description: &movie.Description,
		Year:        &movie.Year,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoriesModel []*model.Category
	for _, category := range categories {
		categoriesModel = append(categoriesModel, &model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: &category.Description,
		})
	}
	return categoriesModel, nil
}

// Movies is the resolver for the movies field.
func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	movies, err := r.MovieDB.FindAll()
	if err != nil {
		return nil, err
	}
	var movieModel []*model.Movie
	for _, movie := range movies {
		movieModel = append(movieModel, &model.Movie{
			ID:          movie.ID,
			Name:        movie.Name,
			Description: &movie.Description,
			Year:        &movie.Year,
		})
	}
	return movieModel, nil
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Movie returns MovieResolver implementation.
func (r *Resolver) Movie() MovieResolver { return &movieResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type movieResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
