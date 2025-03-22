package api_v1

import (
	"fmt"
	"net/http"
	"oopLab1/config"
	"oopLab1/core/customer"
	"oopLab1/pkg/jwt"
	"oopLab1/pkg/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var customerService = customer.NewCustomerService(config.GetConfig().Database)

func Login(ctx echo.Context) error {
	customer := new(customer.Customer)

	if err := ctx.Bind(customer); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	customer.ID = uuid.New().String()
	logger.Info(fmt.Sprintf("Created customer with id %s", customer.ID))

	err := customerService.CreateCustomer(customer)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	token, err := jwt.GenerateJWT(customer.ID, "")

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func GetCustomer(ctx echo.Context) error {
	id := ctx.Param("id")
	customer, err := customerService.GetCustomer(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, customer)
}
