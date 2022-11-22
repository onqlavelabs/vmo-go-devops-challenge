package server

import (
    "github.com/labstack/echo/v4"
)

type EchoServer struct {
    Address string
}

func (s *EchoServer) Serve() {
    server := echo.New()

    server.Logger.Fatal(server.Start(s.Address))
}
