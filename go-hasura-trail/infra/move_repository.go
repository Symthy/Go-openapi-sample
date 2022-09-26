package infra

import (
	"fmt"

	"github.com/Symthy/golang-practices/go-hasura-trail/graph/model"
	"github.com/go-pg/pg"
)

type MovieRepository struct {
	db *pg.DB
}

func NewMovieRepository(db *pg.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r MovieRepository) CreateMovie(input model.NewMovie) (*model.Movie, error) {
	movie := model.Movie{
		Title: input.Title,
		URL:   input.URL,
	}

	_, err := r.db.Model(&movie).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting new movie: %v", err)
	}

	return &movie, nil
}

func (r MovieRepository) FindAll() ([]*model.Movie, error) {
	var movies []*model.Movie

	err := r.db.Model(&movies).Select()
	if err != nil {
		return nil, err
	}

	return movies, nil
}
