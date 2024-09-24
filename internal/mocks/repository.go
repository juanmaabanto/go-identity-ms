package mocks

import (
	"context"

	genrepo "github.com/sofisoft-tech/go-common/gen-repo"
)

type MockRepository[D genrepo.Document] struct {
	StubFindById  func(ctx context.Context, id string) (*D, error)
	StubFindOne   func(ctx context.Context, filters map[string]any) (*D, error)
	StubInsertOne func(ctx context.Context, document D) (string, error)
	StubUpdateOne func(ctx context.Context, id string, document D) error
}

func (mr MockRepository[D]) FindById(ctx context.Context, id string) (*D, error) {
	return mr.StubFindById(ctx, id)
}

func (mr MockRepository[D]) FindOne(ctx context.Context, filters map[string]any) (*D, error) {
	return mr.StubFindOne(ctx, filters)
}

func (mr MockRepository[D]) InsertOne(ctx context.Context, document D) (string, error) {
	return mr.StubInsertOne(ctx, document)
}

func (mr MockRepository[D]) UpdateOne(ctx context.Context, id string, document D) error {
	return mr.StubUpdateOne(ctx, id, document)
}
