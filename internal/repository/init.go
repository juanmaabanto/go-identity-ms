package repository

import (
	"github.com/sofisoft-tech/go-common/database"
	"github.com/sofisoft-tech/go-identity-ms/internal/domain/repotypes"
	"github.com/sofisoft-tech/go-identity-ms/internal/repository/clientapp"
	"github.com/sofisoft-tech/go-identity-ms/internal/repository/user"
)

type Repositories struct {
	ClientApp repotypes.ClientAppRepository
	User      repotypes.UserRepository
}

func InitRepositories(db database.Database) Repositories {
	return Repositories{
		ClientApp: clientapp.New(db),
		User:      user.New(db),
	}
}
