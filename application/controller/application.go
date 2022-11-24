package controller

import (
    "context"
    "fmt"
    "net/http"

    "github.com/dinhtp/vmo-go-devops-challenge/application/domain/application"
    "github.com/dinhtp/vmo-go-devops-challenge/application/domain/auth"
    "github.com/dinhtp/vmo-go-devops-challenge/application/message"
    "github.com/dinhtp/vmo-go-devops-challenge/application/model"
    "github.com/dinhtp/vmo-go-devops-challenge/application/util"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "go.mongodb.org/mongo-driver/mongo"
)

type ApplicationController struct {
    server *echo.Echo
    db     *mongo.Client
}

func NewApplicationController(server *echo.Echo, db *mongo.Client) *ApplicationController {
    return &ApplicationController{db: db, server: server}
}

func (c *ApplicationController) RegisterHandler() {
    group := c.server.Group(fmt.Sprintf("/%s/%s", message.BasePath, message.ApplicationsPath))

    jwtConfig := middleware.JWTConfig{Claims: &model.JwtAuthClaim{}, SigningKey: []byte(auth.DefaultTokenSecret)}
    group.Use(middleware.JWTWithConfig(jwtConfig))

    group.GET("/:id", c.Get)
    group.GET("", c.List)
    group.POST("", c.Create)
    group.PUT("/:id", c.Update)
    group.DELETE("/:id", c.Delete)
}

// Get godoc
// @Summary      Get an application detail
// @Description  Get an application by ID
// @Tags         applications
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Application ID"
// @Success      200  {object}  message.Application
// @Failure      500  {object}  interface{} "{"message":"error_description"}"
// @Router       /applications/{id} [get]
func (c *ApplicationController) Get(e echo.Context) error {
    request := &message.OneApplicationRequest{ID: e.Param("id")}

    err := request.Validate()
    if err != nil {
        return echo.NewHTTPError(http.StatusPreconditionFailed, err.Error())
    }

    result, err := application.NewService(application.NewRepository(c.db)).Get(context.Background(), request)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}

// List godoc
// @Summary      Get applications list
// @Description  Get applications list by page and limit, filter by "type", "name" and "enabled"
// @Tags         applications
// @Accept       json
// @Produce      json
// @Param        page      query		int 	false	"current page"
// @Param        limit     query		int 	false	"page limit"
// @Param        type      query		string	false	"current page"
// @Param        name      query		string	false	"page limit"
// @Param        enabled   query        string	false	"page limit"
// @Success      200  {object}  message.ListApplicationResponse
// @Failure      500  {object}  interface{} "{"message":"error_description"}"
// @Router       /applications [get]
func (c *ApplicationController) List(e echo.Context) error {
    request := &message.ListApplicationRequest{
        Page:    uint(util.StringToInt(e.QueryParam("page"))),
        Limit:   uint(util.StringToInt(e.QueryParam("limit"))),
        Type:    e.QueryParam("type"),
        Name:    e.QueryParam("name"),
        Enabled: e.QueryParam("enabled"),
    }

    err := request.Validate()
    if err != nil {
        return echo.NewHTTPError(http.StatusPreconditionFailed, err.Error())
    }

    result, err := application.NewService(application.NewRepository(c.db)).List(context.Background(), request)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}

// Create godoc
// @Summary      Create new application
// @Description  Create new application
// @Tags         applications
// @Accept       json
// @Produce      json
// @Param        application    body   message.Application true  "Application Data"
// @Success      204  {object}  message.Application
// @Failure      500  {object}  interface{} "{"message":"error_description"}"
// @Router       /applications [post]
func (c *ApplicationController) Create(e echo.Context) error {
    payload := new(message.Application)
    if err := e.Bind(payload); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    err := payload.Validate()
    if err != nil {
        return echo.NewHTTPError(http.StatusPreconditionFailed, err.Error())
    }

    result, err := application.NewService(application.NewRepository(c.db)).Create(context.Background(), payload)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return e.JSONPretty(http.StatusCreated, result, "  ")
}

// Update godoc
// @Summary      Update application detail
// @Description  Update application detail by ID
// @Tags         applications
// @Accept       json
// @Produce      json
// @Param        application body      message.Application true  "Application Data"
// @Param        id          path      string              true  "Application ID"
// @Success      200  {object}  message.Application
// @Failure      500  {object}  interface{} "{"message":"error_description"}"
// @Router       /applications/{id} [put]
func (c *ApplicationController) Update(e echo.Context) error {
    payload := new(message.Application)
    if err := e.Bind(payload); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    err := payload.Validate()
    if err != nil {
        return echo.NewHTTPError(http.StatusPreconditionFailed, err.Error())
    }

    payload.ID = e.Param("id")
    result, err := application.NewService(application.NewRepository(c.db)).Update(context.Background(), payload)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return e.JSONPretty(http.StatusOK, result, "  ")
}

// Delete godoc
// @Summary      Delete an application
// @Description  Delete an application detail by ID
// @Tags         applications
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Application ID"
// @Success      200  {object}  interface{} "{}"
// @Failure      500  {object}  interface{} "{"message":"error_description"}"
// @Router       /applications/{id} [delete]
func (c *ApplicationController) Delete(e echo.Context) error {
    request := &message.OneApplicationRequest{ID: e.Param("id")}

    err := request.Validate()
    if err != nil {
        return echo.NewHTTPError(http.StatusPreconditionFailed, err.Error())
    }

    err = application.NewService(application.NewRepository(c.db)).Delete(context.Background(), request)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return e.NoContent(http.StatusNoContent)
}
