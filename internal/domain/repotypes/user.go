package repotypes

import (
	"context"

	genrepo "github.com/sofisoft-tech/go-common/gen-repo"
	"github.com/sofisoft-tech/go-identity-ms/internal/domain/model"
)

type UserRepository interface {
	genrepo.Repository[model.User]
	AccessFailed(context.Context, *model.User) error
	FindByEmail(context.Context, string) (*model.User, error)
	FindByUserName(context.Context, string) (*model.User, error)
	ResetAccessFailedCount(context.Context, *model.User) error
}
