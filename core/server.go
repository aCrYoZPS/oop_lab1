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

	customerGroup := es.server.Group("/customers")

	customerGroup.POST("/register", api_v1.RegisterCustomer)
	customerGroup.POST("/login", api_v1.LoginCustomer)
	customerGroup.GET("/:id", api_v1.GetCustomer, auth.JWTMiddleware())
	customerGroup.DELETE("/:id", api_v1.DeleteCustomer, auth.JWTMiddleware())
	customerGroup.PATCH("/:id", api_v1.UpdateCustomer, auth.JWTMiddleware())

	companyGroup := es.server.Group("/company")

	companyGroup.POST("/register", api_v1.RegisterCompany)
	companyGroup.POST("/login", api_v1.LoginCompany)
	companyGroup.GET("/:id", api_v1.GetCompany, auth.JWTMiddleware())
	companyGroup.DELETE("/:id", api_v1.DeleteCompany, auth.JWTMiddleware())
	companyGroup.PATCH("/:id", api_v1.UpdateCompany, auth.JWTMiddleware())

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
