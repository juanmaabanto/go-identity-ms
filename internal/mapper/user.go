package mapper

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	identityv1 "github.com/sofisoft-tech/go-contracts/gen/go/identity/v1"

	"github.com/sofisoft-tech/go-identity-ms/internal/domain/model"
)

func GetCreateUserRequest(input *identityv1.CreateUserRequest) model.User {
	return model.User{
		WorkspaceId:            input.WorkspaceId,
		UserName:               input.UserName,
		Email:                  input.Email,
		NormalizedUserName:     strings.ToUpper(input.UserName),
		NormalizedEmail:        strings.ToUpper(input.Email),
		FirstName:              input.FirstName,
		LastName:               input.LastName,
		Alias:                  input.Alias,
		LockoutEnabled:         input.LockoutEnabled,
		PasswordExpiresEnabled: input.PasswordExpiresEnabled,
		RequestPasswordChange:  input.RequestPasswordChange,
		SecurityStamp:          uuid.New().String(),
		Active:                 true,
		CreatedAt:              time.Now().UTC().Unix(),
		CreatedBy:              input.UserName,
	}
}

func GetUserResponse(user *model.User) *identityv1.UserResponse {
	return &identityv1.UserResponse{
		Id:                     user.ID,
		WorkspaceId:            user.WorkspaceId,
		UserName:               user.UserName,
		Email:                  user.Email,
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		Alias:                  user.Alias,
		ImagerUri:              user.ImageUri,
		LockoutEnabled:         user.LockoutEnabled,
		PasswordExpiresEnabled: user.PasswordExpiresEnabled,
		RequestPasswordChange:  user.RequestPasswordChange,
		SecurityStamp:          user.SecurityStamp,
		Active:                 user.Active,
		CreatedAt:              timestamppb.New(time.Unix(user.CreatedAt, 0)),
		CreatedBy:              user.CreatedBy,
	}
}

func GetSignInResponse(user *model.User) *identityv1.SignInResponse {
	var expireAt *timestamppb.Timestamp

	if user.PasswordExpires != nil {
		expireAt = timestamppb.New(time.Unix(*user.PasswordExpires, 0))
	}
	return &identityv1.SignInResponse{
		UserId:                 user.ID,
		UserName:               user.UserName,
		SecurityStamp:          user.SecurityStamp,
		PasswordExpiresEnabled: user.PasswordExpiresEnabled,
		PasswordExpires:        expireAt,
		RequestPasswordChange:  user.RequestPasswordChange,
	}
}
