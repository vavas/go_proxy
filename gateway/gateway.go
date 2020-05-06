package gateway

import (
	"bitbucket.org/telemetryapp/go_rabbitmq/config"
	"bitbucket.org/telemetryapp/go_rabbitmq/server"
	"log"
)

type Gateway struct {
	server *server.Server
	config *config.Config
}

func New(config *config.Config) (*Gateway, error)  {
	gateway := &Gateway{
		config: config,
	}

	svr, err := server.New(config)
	if err != nil {
		log.Println(err)
	}

	gateway.server = svr
	log.Println("Server created")

	return gateway, nil


	return gateway, nil
}

