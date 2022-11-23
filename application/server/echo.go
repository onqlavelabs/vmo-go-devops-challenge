package server

import (
    "github.com/dinhtp/vmo-go-devops-challenge/application/controller"
    "github.com/labstack/echo/v4"
    "go.mongodb.org/mongo-driver/mongo"
)

type EchoServer struct {
    db      *mongo.Client
    address string
}

func NewEchoServer(address string, db *mongo.Client) *EchoServer {
    return &EchoServer{
        db:      db,
        address: address,
    }
}

func (s *EchoServer) Serve() {
    server := echo.New()

    controller.InitSwagger(server)
    controller.NewAuthController(server).RegisterHandler()
    controller.NewApplicationController(server, s.db).RegisterHandler()

    server.Logger.Fatal(server.Start(s.address))
}
