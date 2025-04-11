package api_v1

import (
	"fmt"
	"net/http"
	"oopLab1/auth"
	"oopLab1/config"
	"oopLab1/core/customer"
	myJWT "oopLab1/pkg/jwt"
	"oopLab1/pkg/logger"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var custService = customer.NewCustomerService(config.GetConfig().Database)

func RegisterCustomer(ctx echo.Context) error {
	cust := new(customer.Customer)

	if err := ctx.Bind(cust); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if !customer.IsValid(cust) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid customer data",
		})
	}

	cust.ID = uuid.New().String()
	logger.Info(fmt.Sprintf("Created cust with id %s", cust.ID))

	token, err := myJWT.GenerateJWT(cust.ID, "cust")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	err = custService.CreateCustomer(cust)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"token": token,
		"id":    cust.ID,
	})
}

func LoginCustomer(ctx echo.Context) error {
	loginRequest := new(auth.LoginRequest)

	if err := ctx.Bind(loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid login request",
		})
	}

	cust, err := custService.GetCustomerByEmail(loginRequest.Username)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "Customer with such username doesn't exist",
		})
	}

	if cust.Password != loginRequest.Password {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Wrong password",
		})
	}

	token, err := myJWT.GenerateJWT(cust.ID, "cust")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
		"id":    cust.ID,
	})
}

func GetCustomer(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if userID != id && role == "cust" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other cust banking account is prohibited",
		})
	}

	cust, err := custService.GetCustomerByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, cust)
}

func DeleteCustomer(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if userID != id && role == "cust" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other cust banking account is prohibited",
		})
	}

	err := custService.DeleteCustomer(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return ctx.NoContent(http.StatusNoContent)
}

func UpdateCustomer(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if userID != id && role == "cust" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other cust banking account is prohibited",
		})
	}

	updatedCustomer := new(customer.Customer)

	if err := ctx.Bind(updatedCustomer); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	cust, err := custService.GetCustomerByID(userID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	customer.UpdateCustomerInfo(cust, updatedCustomer)

	if !customer.IsValid(updatedCustomer) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid customer data",
		})
	}

	err = custService.UpdateCustomer(updatedCustomer)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updatedCustomer)
}
