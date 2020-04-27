package graph

import "edwardlie-graphql-server/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	movies []*model.Movie
}
