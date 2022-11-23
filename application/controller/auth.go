package controller

import (
    "context"
    "fmt"
    "net/http"

    "github.com/dinhtp/vmo-go-devops-challenge/application/domain/auth"
    "github.com/dinhtp/vmo-go-devops-challenge/application/message"
    "github.com/labstack/echo/v4"
)

type AuthController struct {
    server *echo.Echo
}

func NewAuthController(server *echo.Echo) *AuthController {
    return &AuthController{server: server}
}

func (c *AuthController) RegisterHandler() {
    group := c.server.Group(fmt.Sprintf("/%s/%s", message.BasePath, message.AuthenticatePath))

    group.POST("", c.Authenticate)
}

// Authenticate godoc
// @Summary      Authenticate existing user
// @Description  Authenticate existing user
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param        application    body   message.AuthRequest true  "Authenticate Data"
// @Success      200  {object}  message.AuthResponse
// @Failure      500  {object}  interface{} "{"message":"error_description"}"
// @Failure      403  {object}  interface{} "{"message":"error_description"}"
// @Router       /authenticate [post]
func (c *AuthController) Authenticate(e echo.Context) error {
    payload := new(message.AuthRequest)
    if err := e.Bind(payload); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    err := payload.Validate()
    if err != nil {
        return echo.NewHTTPError(http.StatusPreconditionFailed, err.Error())
    }

    result, err := auth.NewService().Authenticate(context.Background(), payload)
    if err != nil {
        return echo.NewHTTPError(http.StatusForbidden, err.Error())
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}
