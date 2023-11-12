package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhodinemma/go-echo-pg/internal/database"
)

type Server interface {
	start() error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}

	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal("server shutdown occurred: %s", err)
		return err
	}

	return nil
}

func (s *EchoServer) registerRoutes() {

}
