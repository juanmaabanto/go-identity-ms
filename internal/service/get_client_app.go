package service

import (
	"context"

	"github.com/sofisoft-tech/go-common/log"
	commonv1 "github.com/sofisoft-tech/go-contracts/gen/go/common/v1"
	identityv1 "github.com/sofisoft-tech/go-contracts/gen/go/identity/v1"

	"github.com/sofisoft-tech/go-identity-ms/internal/mapper"
	"github.com/sofisoft-tech/go-identity-ms/pkg/constant"
)

func (svc service) GetClientApp(ctx context.Context, input *identityv1.GetClientAppRequest) (response *commonv1.Response, err error) {
	response = &commonv1.Response{}

	if input == nil || len(input.ClientId) == 0 {
		response.Code, response.Message = constant.ErrMissingFields.Message(ctx)
		return response, err
	}

	clientApp, err := svc.deps.Repo.ClientApp.FindByClientId(ctx, input.ClientId)
	if err != nil {
		return response, err
	}

	if clientApp == nil {
		response.Code, response.Message = constant.ErrClientAppNotFound.Message(ctx)
		return response, err
	}

	output := mapper.GetClientAppResponse(clientApp)

	dataBytes, err := svc.deps.Proto.Marshal(output)
	if err != nil {
		log.Error(err.Error(), log.String("clientId", input.ClientId))
		return response, err
	}

	response.Message = "Ok"
	response.Data = dataBytes

	return response, nil
}
