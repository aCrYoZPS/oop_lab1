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
	companyGroup.GET("", api_v1.GetAllCompanies, auth.JWTMiddleware())
	companyGroup.DELETE("/:id", api_v1.DeleteCompany, auth.JWTMiddleware())
	companyGroup.PATCH("/:id", api_v1.UpdateCompany, auth.JWTMiddleware())

	accountGroup := es.server.Group("/account")
	accountGroup.POST("", api_v1.CreateAccount, auth.JWTMiddleware())
	accountGroup.DELETE("/:acc_id", api_v1.DeleteAccount, auth.JWTMiddleware())
	accountGroup.GET("/:acc_id", api_v1.GetAccount, auth.JWTMiddleware())
	accountGroup.GET("/all", api_v1.GetAllAccounts, auth.JWTMiddleware())
	accountGroup.GET("", api_v1.GetAllAccountsByOwner, auth.JWTMiddleware())
	accountGroup.POST("/:acc_id/withdaw_topup", api_v1.CreateSelfTransaction, auth.JWTMiddleware())
	accountGroup.POST("/:acc_id/transfer/:target_id", api_v1.CreateTransaction, auth.JWTMiddleware())

	staffGroup := es.server.Group("/staff")
	staffGroup.POST("/register", api_v1.RegisterStaffMember)
	staffGroup.POST("/login", api_v1.LoginStaffMember)
	staffGroup.GET("/:id", api_v1.GetStaffMember, auth.JWTMiddleware())
	staffGroup.DELETE("/:id", api_v1.DeleteStaffMember, auth.JWTMiddleware())
	staffGroup.PATCH("/:id", api_v1.UpdateStaffMember, auth.JWTMiddleware())
	staffGroup.POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc)

	bankGroup := es.server.Group("/bank")
	bankGroup.POST("/register", api_v1.RegisterBank)
	bankGroup.GET("/:id", api_v1.GetBankByID, auth.JWTMiddleware())
	bankGroup.GET("", api_v1.GetAllBanks, auth.JWTMiddleware())
	bankGroup.DELETE("/:id", api_v1.DeleteBank, auth.JWTMiddleware())
	bankGroup.PATCH("/:id", api_v1.UpdateBank, auth.JWTMiddleware())

	loanGroup := es.server.Group("/loans")
	loanGroup.POST("", api_v1.CreateLoan, auth.JWTMiddleware())
	loanGroup.DELETE("/:id", api_v1.DeleteLoan, auth.JWTMiddleware())
	loanGroup.GET("/:id", api_v1.GetLoan, auth.JWTMiddleware())
	loanGroup.GET("/all", api_v1.GetAllLoans, auth.JWTMiddleware())
	loanGroup.GET("", api_v1.GetAllLoansByOwner, auth.JWTMiddleware())
	loanGroup.PATCH("/:id", api_v1.UpdateLoan, auth.JWTMiddleware())

	salaryGroup := es.server.Group("/salary")
	salaryGroup.POST("", api_v1.CreateSalaryProject, auth.JWTMiddleware())
	salaryGroup.DELETE("/:id", api_v1.DeleteSalaryProject, auth.JWTMiddleware())
	salaryGroup.GET("/:id", api_v1.GetSalaryProject, auth.JWTMiddleware())
	salaryGroup.GET("/all", api_v1.GetAllSalaryProjects, auth.JWTMiddleware())
	salaryGroup.GET("", api_v1.GetAllSalaryProjectsByCompany, auth.JWTMiddleware())
	salaryGroup.PATCH("/:id", api_v1.UpdateSalaryProject, auth.JWTMiddleware())

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
