package core

import (
	"fmt"
	api_v1 "oopLab1/api/v1"
	"oopLab1/auth"
	"oopLab1/config"
	"oopLab1/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Start()
}

type EchoServer struct {
	server *echo.Echo
}

func NewEchoServer() Server {
	var es EchoServer

	es.server = echo.New()

	es.server.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.Info(fmt.Sprintf("REQUEST: uri: %v, status: %v\n", v.URI, v.Status))
			} else {
				logger.Error(fmt.Sprintf("REQUEST: uri: %v, status: %v, err: %v\n", v.URI, v.Status, v.Error))
			}
			return nil
		},
	}))

	es.server.POST("/customers/register", api_v1.Register)
	es.server.POST("/customers/login", api_v1.Login)
	es.server.GET("/customers/:id", api_v1.GetCustomer, auth.JWTMiddleware())
	es.server.DELETE("/customers/:id", api_v1.DeleteCustomer, auth.JWTMiddleware())
	es.server.PATCH("/customers/:id", api_v1.UpdateCustomer, auth.JWTMiddleware())

	return &es
}

func (es *EchoServer) Start() {
	configuration := config.GetConfig()
	serverConfig := configuration.Server
	addres := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	if err := es.server.Start(addres); err != nil {
		logger.Fatal("Server has failed to start")
	}
}
