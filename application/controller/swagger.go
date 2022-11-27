package controller

import (
    _ "github.com/dinhtp/vmo-go-devops-challenge/application/docs"
    "github.com/labstack/echo/v4"
    echoSwagger "github.com/swaggo/echo-swagger"
)

func InitSwagger(echo *echo.Echo) {
    echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
