package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/sofisoft-tech/go-common/log"
	commonv1 "github.com/sofisoft-tech/go-contracts/gen/go/common/v1"
	identityv1 "github.com/sofisoft-tech/go-contracts/gen/go/identity/v1"
	"github.com/sofisoft-tech/go-identity-ms/internal/mapper"
	"github.com/sofisoft-tech/go-identity-ms/pkg/constant"
)

func (svc service) CreateClientApp(ctx context.Context, input *identityv1.CreateClientAppRequest) (response *commonv1.Response, err error) {
	response = &commonv1.Response{}

	if !svc.deps.Validator.ValidateClientAppFields(input) {
		response.Code, response.Message = constant.ErrMissingFields.Message(ctx)
		return response, err
	}

	clientApp := mapper.GetCreateClientAppRequest(input)

	// generate clientId and clientSecret
	bytes := make([]byte, 32)
	_, err = rand.Read(bytes)
	if err != nil {
		log.Error(err.Error())
		return response, err
	}

	clientApp.ClientId = uuid.New().String()
	clientApp.ClientSecret = hex.EncodeToString(bytes)

	// insert clientApp
	clientApp.ID, err = svc.deps.Repo.ClientApp.InsertOne(ctx, clientApp)
	if err != nil {
		log.Error(err.Error())
		return response, err
	}

	output := mapper.GetClientAppResponse(&clientApp)

	dataBytes, err := svc.deps.Proto.Marshal(output)
	if err != nil {
		log.Error(err.Error())
		return response, err
	}

	response.Message = "Ok"
	response.Data = dataBytes

	return response, nil
}
