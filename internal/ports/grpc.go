package ports

import (
	"github.com/dapr/go-sdk/service/common"

	daprs "github.com/sofisoft-tech/go-common/dapr-server"
	identityv1 "github.com/sofisoft-tech/go-contracts/gen/go/identity/v1"

	"github.com/sofisoft-tech/go-identity-ms/internal/service"
)

func GetHandlers(svc service.Service) map[string]common.ServiceInvocationHandler {
	return map[string]common.ServiceInvocationHandler{
		"authenticate":      daprs.InvocationHandlerWithContent(&identityv1.SignInRequest{}, svc.Authenticate),
		"create-client-app": daprs.InvocationHandlerWithContent(&identityv1.CreateClientAppRequest{}, svc.CreateClientApp),
		"create-user":       daprs.InvocationHandlerWithContent(&identityv1.CreateUserRequest{}, svc.CreateUser),
		"get-client-app":    daprs.InvocationHandlerWithContent(&identityv1.GetClientAppRequest{}, svc.GetClientApp),
	}
}
