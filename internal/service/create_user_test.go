package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	commonv1 "github.com/sofisoft-tech/go-contracts/gen/go/common/v1"
	identityv1 "github.com/sofisoft-tech/go-contracts/gen/go/identity/v1"
	"github.com/sofisoft-tech/go-identity-ms/internal/domain/model"
	"github.com/sofisoft-tech/go-identity-ms/internal/mocks"
	"github.com/sofisoft-tech/go-identity-ms/internal/repository"
	"github.com/sofisoft-tech/go-identity-ms/internal/validation"
)

func TestServiceCreateUser(t *testing.T) {
	inputValid := &identityv1.CreateUserRequest{
		Email:       "test@fake.com",
		UserName:    "fakename",
		FirstName:   "fake",
		LastName:    "fake",
		Password:    "123456",
		WorkspaceId: "123",
	}
	type fields struct {
		deps ServiceDeps
	}
	type args struct {
		ctx   context.Context
		input *identityv1.CreateUserRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse *commonv1.Response
		wantErr      bool
	}{
		{
			name: "empty request",
			fields: fields{
				deps: ServiceDeps{
					Validator: validation.New(),
				},
			},
			args: args{
				ctx:   context.Background(),
				input: nil,
			},
			wantResponse: &commonv1.Response{
				Code:    "ERROR_MISSING_FIELDS",
				Message: "Required fields not supplied",
			},
			wantErr: false,
		},
		{
			name: "validate fields",
			fields: fields{
				deps: ServiceDeps{
					Validator: validation.New(),
				},
			},
			args: args{
				ctx:   context.Background(),
				input: &identityv1.CreateUserRequest{},
			},
			wantResponse: &commonv1.Response{
				Code:    "ERROR_MISSING_FIELDS",
				Message: "Required fields not supplied",
			},
			wantErr: false,
		},
		{
			name: "error find username ",
			fields: fields{
				deps: ServiceDeps{
					Repo: repository.Repositories{
						User: MockUserRepository{
							StubFindByUserName: ErrorFindByUserName,
						},
					},
					Validator: validation.New(),
				},
			},
			args: args{
				ctx:   context.Background(),
				input: inputValid,
			},
			wantResponse: &commonv1.Response{},
			wantErr:      true,
		},
		{
			name: "username exists",
			fields: fields{
				deps: ServiceDeps{
					Repo: repository.Repositories{
						User: MockUserRepository{
							StubFindByUserName: OkFindByUserName,
						},
					},
					Validator: validation.New(),
				},
			},
			args: args{
				ctx:   context.Background(),
				input: inputValid,
			},
			wantResponse: &commonv1.Response{
				Code:    "ERROR_USER_NAME_EXISTS",
				Message: "The username is already being used in another account",
			},
			wantErr: false,
		},
		{
			name: "error find by email ",
			fields: fields{
				deps: ServiceDeps{
					Repo: repository.Repositories{
						User: MockUserRepository{
							StubFindByUserName: NotFoundUserName,
							StubFindByEmail:    ErrorFindByEmail,
						},
					},
					Validator: validation.New(),
				},
			},
			args: args{
				ctx:   context.Background(),
				input: inputValid,
			},
			wantResponse: &commonv1.Response{},
			wantErr:      true,
		},
		{
			name: "email exists",
			fields: fields{
				deps: ServiceDeps{
					Repo: repository.Repositories{
						User: MockUserRepository{
							StubFindByUserName: NotFoundEmail,
							StubFindByEmail:    OkFindByEmail,
						},
					},
					Validator: validation.New(),
				},
			},
			args: args{
				ctx:   context.Background(),
				input: inputValid,
			},
			wantResponse: &commonv1.Response{
				Code:    "ERROR_EMAIL_EXISTS",
				Message: "The email is already being used in another account",
			},
			wantErr: false,
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			svc := service{
				deps: tt.fields.deps,
			}
			gotResponse, err := svc.CreateUser(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("service.CreateUser() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

type MockUserRepository struct {
	mocks.MockRepository[model.User]
	StubAccessFailed           func(ctx context.Context, user *model.User) error
	StubFindByEmail            func(ctx context.Context, email string) (*model.User, error)
	StubFindByUserName         func(ctx context.Context, userName string) (*model.User, error)
	StubResetAccessFailedCount func(ctx context.Context, user *model.User) error
}

func (mr MockUserRepository) AccessFailed(ctx context.Context, user *model.User) error {
	return mr.StubAccessFailed(ctx, user)
}

func (mr MockUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return mr.StubFindByEmail(ctx, email)
}

func (mr MockUserRepository) FindByUserName(ctx context.Context, userName string) (*model.User, error) {
	return mr.StubFindByUserName(ctx, userName)
}

func (mr MockUserRepository) ResetAccessFailedCount(ctx context.Context, user *model.User) error {
	return mr.StubResetAccessFailedCount(ctx, user)
}

var (
	ErrorFindByEmail = func(ctx context.Context, email string) (*model.User, error) {
		return nil, errors.New("fake error find by email")
	}
	ErrorFindByUserName = func(ctx context.Context, userName string) (*model.User, error) {
		return nil, errors.New("fake error find by username")
	}
	NotFoundUserName = func(ctx context.Context, userName string) (*model.User, error) {
		return nil, nil
	}
	NotFoundEmail = func(ctx context.Context, email string) (*model.User, error) {
		return nil, nil
	}
	OkFindByEmail = func(ctx context.Context, email string) (*model.User, error) {
		return &model.User{}, nil
	}
	OkFindByUserName = func(ctx context.Context, userName string) (*model.User, error) {
		return &model.User{}, nil
	}
)
