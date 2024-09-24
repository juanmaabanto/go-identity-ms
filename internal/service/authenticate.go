package service

import (
	"context"

	"github.com/sofisoft-tech/go-common/crypto"
	commonv1 "github.com/sofisoft-tech/go-contracts/gen/go/common/v1"
	identityv1 "github.com/sofisoft-tech/go-contracts/gen/go/identity/v1"
	"github.com/sofisoft-tech/go-identity-ms/internal/mapper"
	"github.com/sofisoft-tech/go-identity-ms/pkg/constant"
)

func (svc service) Authenticate(ctx context.Context, input *identityv1.SignInRequest) (response *commonv1.Response, err error) {
	response = &commonv1.Response{}

	if len(input.UserName) == 0 {
		response.Code, response.Message = constant.ErrMissingFields.Message(ctx)
		return response, err
	}

	user, err := svc.deps.Repo.User.FindByUserName(ctx, input.UserName)
	if err != nil {
		return response, err
	}

	if user == nil {
		response.Code, response.Message = constant.ErrUserNotFound.Message(ctx)
		return response, err
	}

	if !user.Active {
		response.Code, response.Message = constant.ErrUserInactive.Message(ctx)
		return response, err
	}

	if user.IsLockedOut() {
		response.Code, response.Message = constant.ErrUserIsLockOut.Message(ctx)
		return response, err
	}

	ok := crypto.CheckPasswordHash(input.Password, user.PasswordHash)

	if ok {
		err = svc.deps.Repo.User.ResetAccessFailedCount(ctx, user)
		if err != nil {
			return response, err
		}

	} else {
		err = svc.deps.Repo.User.AccessFailed(ctx, user)
		if err != nil {
			return response, err
		}

		if user.LockoutEnd == nil {
			response.Code, response.Message = constant.ErrWrongPassword.Message(ctx)
		} else {
			response.Code, response.Message = constant.ErrWrongPasswordLocked.Message(ctx)
		}

		return response, err
	}

	output := mapper.GetSignInResponse(user)

	dataBytes, err := svc.deps.Proto.Marshal(output)
	if err != nil {
		return response, err
	}

	response.Message = "Ok"
	response.Data = dataBytes

	return response, nil
}
