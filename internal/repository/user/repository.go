package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/sofisoft-tech/go-common/database"
	genrepo "github.com/sofisoft-tech/go-common/gen-repo"
	"github.com/sofisoft-tech/go-common/log"
	"github.com/sofisoft-tech/go-identity-ms/internal/domain/model"
	"github.com/sofisoft-tech/go-identity-ms/internal/domain/repotypes"
)

type repository struct {
	genrepo.GRepository[model.User]
}

func New(db database.Database) repotypes.UserRepository {
	return repository{
		GRepository: genrepo.GRepository[model.User]{
			Database: db,
		},
	}
}

func (r repository) AccessFailed(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	if !user.LockoutEnabled {
		return nil
	}

	user.AccessFailedCount += 1

	if user.AccessFailedCount >= 3 {
		lockoutEnd := time.Now().Add(time.Minute * 5).Unix()
		user.LockoutEnd = &lockoutEnd
	} else {
		user.LockoutEnd = nil
	}

	err := r.UpdateOne(ctx, user.ID, *user)
	if err != nil {
		log.Error(err.Error(), log.String("username", user.UserName))
		return err
	}
	return nil
}

func (r repository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	filters := make(map[string]any)
	filters["normalizedEmail"] = strings.ToUpper(email)

	user, err := r.FindOne(ctx, filters)
	if err != nil {
		log.Error(err.Error(), log.String("email", email))
		return nil, err
	}

	return user, nil
}

func (r repository) FindByUserName(ctx context.Context, userName string) (*model.User, error) {
	filters := make(map[string]any)
	filters["normalizedUserName"] = strings.ToUpper(userName)

	user, err := r.FindOne(ctx, filters)
	if err != nil {
		log.Error(err.Error(), log.String("username", userName))
		return nil, err
	}

	return user, nil
}

func (r repository) ResetAccessFailedCount(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.New("user for reset fail count is nil")
	}

	if user.AccessFailedCount == 0 {
		return nil
	}

	user.AccessFailedCount = 0
	user.LockoutEnd = nil

	err := r.UpdateOne(ctx, user.ID, *user)
	if err != nil {
		log.Error(err.Error(), log.String("username", user.UserName))
		return err
	}
	return nil
}
