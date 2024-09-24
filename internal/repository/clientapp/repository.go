package clientapp

import (
	"context"

	"github.com/sofisoft-tech/go-common/database"
	genrepo "github.com/sofisoft-tech/go-common/gen-repo"
	"github.com/sofisoft-tech/go-common/log"

	"github.com/sofisoft-tech/go-identity-ms/internal/domain/model"
	"github.com/sofisoft-tech/go-identity-ms/internal/domain/repotypes"
)

type repository struct {
	genrepo.GRepository[model.ClientApp]
}

func New(db database.Database) repotypes.ClientAppRepository {
	return repository{
		GRepository: genrepo.GRepository[model.ClientApp]{
			Database: db,
		},
	}
}

func (r repository) FindByClientId(ctx context.Context, clientId string) (*model.ClientApp, error) {
	filters := make(map[string]any)
	filters["clientId"] = clientId

	clientApp, err := r.FindOne(ctx, filters)
	if err != nil {
		log.Error(err.Error(), log.String("clientId", clientId))
		return nil, err
	}

	return clientApp, nil
}
