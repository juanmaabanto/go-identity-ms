package validation

import identityv1 "github.com/sofisoft-tech/go-contracts/gen/go/identity/v1"

type Validator interface {
	ValidateClientAppFields(*identityv1.CreateClientAppRequest) bool
	ValidateUserFields(*identityv1.CreateUserRequest) bool
}

type validator struct {
}

func New() Validator {
	return validator{}
}

func (validator) ValidateClientAppFields(input *identityv1.CreateClientAppRequest) bool {
	if input == nil {
		return false
	}

	if input.Name != "" {
		return true
	} else {
		return false
	}
}

func (validator) ValidateUserFields(user *identityv1.CreateUserRequest) bool {
	if user == nil {
		return false
	}

	fieldsPass := user.Email != "" && user.UserName != "" &&
		user.FirstName != "" && user.LastName != "" &&
		user.Password != "" && user.WorkspaceId != ""

	if fieldsPass {
		return true
	} else {
		return false
	}
}
