package service

import (
	"context"

	"github.com/sofisoft-tech/go-common/protow"
	commonv1 "github.com/sofisoft-tech/go-contracts/gen/go/common/v1"
	identityv1 "github.com/sofisoft-tech/go-contracts/gen/go/identity/v1"

	"github.com/sofisoft-tech/go-identity-ms/internal/repository"
	"github.com/sofisoft-tech/go-identity-ms/internal/validation"
)

type Service interface {
	// clientApp
	CreateClientApp(context.Context, *identityv1.CreateClientAppRequest) (*commonv1.Response, error)
	GetClientApp(context.Context, *identityv1.GetClientAppRequest) (*commonv1.Response, error)
	// user
	Authenticate(context.Context, *identityv1.SignInRequest) (*commonv1.Response, error)
	CreateUser(context.Context, *identityv1.CreateUserRequest) (*commonv1.Response, error)
}

type ServiceDeps struct {
	Proto     protow.Proto
	Repo      repository.Repositories
	Validator validation.Validator
}

type service struct {
	deps ServiceDeps
}

func NewService(deps ServiceDeps) service {
	return service{
		deps: deps,
	}
}
