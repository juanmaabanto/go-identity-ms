package service

import (
	"context"

	"github.com/sofisoft-tech/go-common/log"
	commonv1 "github.com/sofisoft-tech/go-contracts/gen/go/common/v1"
	identityv1 "github.com/sofisoft-tech/go-contracts/gen/go/identity/v1"

	"github.com/sofisoft-tech/go-identity-ms/internal/mapper"
	"github.com/sofisoft-tech/go-identity-ms/pkg/constant"
)

func (svc service) CreateUser(ctx context.Context, input *identityv1.CreateUserRequest) (response *commonv1.Response, err error) {
	response = &commonv1.Response{}

	if !svc.deps.Validator.ValidateUserFields(input) {
		response.Code, response.Message = constant.ErrMissingFields.Message(ctx)
		return response, err
	}

	exists, err := svc.deps.Repo.User.FindByUserName(ctx, input.UserName)
	if err != nil {
		return response, err
	}

	if exists != nil {
		response.Code, response.Message = constant.ErrUserNameExists.Message(ctx)
		return response, err
	}

	exists, err = svc.deps.Repo.User.FindByEmail(ctx, input.Email)
	if err != nil {
		return response, err
	}

	if exists != nil {
		response.Code, response.Message = constant.ErrEmailExists.Message(ctx)
		return response, err
	}

	user := mapper.GetCreateUserRequest(input)
	user.SetPasswordHash(input.Password)

	user.ID, err = svc.deps.Repo.User.InsertOne(ctx, user)
	if err != nil {
		log.Error(err.Error(), log.String("username", user.UserName))
		return response, err
	}

	output := mapper.GetUserResponse(&user)

	dataBytes, err := svc.deps.Proto.Marshal(output)
	if err != nil {
		log.Error(err.Error(), log.String("username", user.UserName))
		return response, err
	}

	response.Message = "Ok"
	response.Data = dataBytes

	return response, nil
}
