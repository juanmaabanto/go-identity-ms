package mapper

import (
	"time"

	identityv1 "github.com/sofisoft-tech/go-contracts/gen/go/identity/v1"
	"github.com/sofisoft-tech/go-identity-ms/internal/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetCreateClientAppRequest(input *identityv1.CreateClientAppRequest) model.ClientApp {
	return model.ClientApp{
		Name:           input.Name,
		Description:    input.Description,
		RedirectUris:   input.RedirectUris,
		AllowedOrigins: input.AllowedOrigins,
		ThirdParty:     input.ThirdParty,
		Active:         true,
		CreatedAt:      time.Now().UTC().Unix(),
		CreatedBy:      input.CreatedBy,
	}
}

func GetClientAppResponse(clientApp *model.ClientApp) *identityv1.ClientAppResponse {
	return &identityv1.ClientAppResponse{
		ClientAppId:    clientApp.ID,
		Name:           clientApp.Name,
		Description:    clientApp.Description,
		ClientId:       clientApp.ClientId,
		ClientSecret:   clientApp.ClientSecret,
		RedirectUris:   clientApp.RedirectUris,
		AllowedOrigins: clientApp.AllowedOrigins,
		ThirdParty:     clientApp.ThirdParty,
		Active:         clientApp.Active,
		CreatedAt:      timestamppb.New(time.Unix(clientApp.CreatedAt, 0)),
		CreatedBy:      clientApp.CreatedBy,
	}
}
