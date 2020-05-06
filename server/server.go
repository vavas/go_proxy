package server

import (
	"bitbucket.org/telemetryapp/go_rabbitmq/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config 		*config.Config
	gin        	*gin.Engine
	httpServer 	*http.Server
	tlsEnabled 	bool
}

func New(config *config.Config) (*Server, error) {

	server := &Server{
		config:     config,
		gin:        gin.New(),
		httpServer: nil,
		tlsEnabled: false,
	}

	server.setupHTTPServer()

	return server, nil
}

// setupHTTPServer is used by New to create and configure an HTTP server based
func (s *Server) setupHTTPServer() error {
	log.Println("Creating HTTP server")
	s.httpServer = &http.Server{
		Addr:              s.config.Server.Listen,
		Handler:           s.gin,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       10 * time.Second,
		// Internal services timeout is 15s, so add 1s for receiving error response from the internal service.
		// Then the gateway can respond with a proper HTTP response.
		WriteTimeout:   16 * time.Second,
	}
	log.Println("HTTP server created")

	return nil
}
