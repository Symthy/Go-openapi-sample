package graph

import (
	"github.com/Symthy/golang-practices/go-hasura-trail/infra"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MovieRepo *infra.MovieRepository
}
