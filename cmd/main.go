package main

import (
	"context"

	"github.com/sofisoft-tech/go-common/arango"
	daprs "github.com/sofisoft-tech/go-common/dapr-server"

	"github.com/sofisoft-tech/go-common/log"
	"github.com/sofisoft-tech/go-common/protow"

	"github.com/sofisoft-tech/go-identity-ms/cmd/config"
	"github.com/sofisoft-tech/go-identity-ms/internal/ports"
	"github.com/sofisoft-tech/go-identity-ms/internal/repository"
	"github.com/sofisoft-tech/go-identity-ms/internal/service"
	"github.com/sofisoft-tech/go-identity-ms/internal/validation"
)

func main() {
	cfg := config.Get()

	db, err := arango.NewDatabase(
		context.Background(),
		cfg.ArangodbEndpoints,
		cfg.ArangodbName,
		cfg.ArangodbUser,
		cfg.ArangodbPassword,
	)
	if err != nil {
		log.Panic(err.Error(), log.String("endpoints", cfg.ArangodbEndpoints), log.String("db", cfg.ArangodbName), log.String("user", cfg.ArangodbUser))
	}

	dbRepositories := repository.InitRepositories(db)

	// Manager Service
	deps := service.ServiceDeps{
		Proto:     protow.NewProto(),
		Repo:      dbRepositories,
		Validator: validation.New(),
	}

	svc := service.NewService(deps)

	// Service Handlers
	opts := daprs.ServerOptions{
		AppPort:         cfg.DaprAppPort,
		ServiceHandlers: ports.GetHandlers(svc),
	}

	server := daprs.NewServer(opts)
	err = server.Start()
	if err != nil {
		log.Panic("failed to start the service: "+err.Error(), log.Int("AppPort", cfg.DaprAppPort))
	}
}
