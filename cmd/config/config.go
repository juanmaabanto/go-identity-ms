package config

import (
	"context"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	daprc "github.com/sofisoft-tech/go-common/dapr-client"
	"github.com/sofisoft-tech/go-common/log"
)

type Config struct {
	// Dapr app config
	DaprAppPort  int `env:"DAPR_APP_PORT" envDefault:"5001"`
	DaprGrpcPort int `env:"DAPR_GRPC_PORT" envDefault:"50001"`

	SecretStoreName string `env:"SECRET_STORE_NAME"`

	// Database
	ArangodbEndpoints string `env:"DB_ENDPOINTS"`
	ArangodbName      string `env:"DB_NAME"`
	ArangodbUser      string
	ArangodbPassword  string
}

func Get() Config {
	cfg := loadEnv()
	client := daprc.NewClient(cfg.DaprGrpcPort,
		daprc.SecretsConfig{
			StoreName: cfg.SecretStoreName,
		},
	)

	if secrets, err := client.GetSecrets(context.Background(), "database/arangodb"); err == nil && len(secrets) > 0 {
		cfg.ArangodbUser = secrets["user"]
		cfg.ArangodbPassword = secrets["password"]
	} else {
		if err != nil {
			log.Panic("failed to read secrets: " + err.Error())
		}
	}

	return cfg
}

func loadEnv() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Warn("missing .env file")
	}

	cfg := Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Panic("failed to parse config: " + err.Error())
	}

	return cfg
}
