package repotypes

import (
	"context"

	genrepo "github.com/sofisoft-tech/go-common/gen-repo"

	"github.com/sofisoft-tech/go-identity-ms/internal/domain/model"
)

type ClientAppRepository interface {
	genrepo.Repository[model.ClientApp]
	FindByClientId(context.Context, string) (*model.ClientApp, error)
}
