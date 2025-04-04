package api_v1

import (
	"fmt"
	"net/http"
	"oopLab1/auth"
	"oopLab1/config"
	"oopLab1/core/company"
	myJWT "oopLab1/pkg/jwt"
	"oopLab1/pkg/logger"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var companyService = company.NewCompanyService(config.GetConfig().Database)

func RegisterCompany(ctx echo.Context) error {
	comp := new(company.Company)

	if err := ctx.Bind(comp); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if !company.IsValid(comp) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid company data",
		})
	}

	comp.ID = uuid.New().String()
	logger.Info(fmt.Sprintf("Created company with id %s", comp.ID))

	token, err := myJWT.GenerateJWT(comp.ID, "company")

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	err = companyService.CreateCompany(comp)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"token": token,
		"id":    comp.ID,
	})
}

func LoginCompany(ctx echo.Context) error {
	var loginRequest = new(auth.LoginRequest)

	if err := ctx.Bind(loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid login request",
		})
	}

	comp, err := companyService.GetCompanyByEmail(loginRequest.Username)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "Company with such username doesn't exist",
		})
	}

	if comp.Password != loginRequest.Password {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Wrong password",
		})
	}

	token, err := myJWT.GenerateJWT(comp.ID, "comp")

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
		"id":    comp.ID,
	})
}

func GetCompany(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if userID != id && role == "company" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other comp banking account is prohibited",
		})
	}

	comp, err := companyService.GetCompanyByID(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, comp)
}

func GetAllCompanies(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if role == "company" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other comp banking account is prohibited",
		})
	}

	companies, err := companyService.GetAllCompanies()

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, companies)
}

func DeleteCompany(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if userID != id && role == "company" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other comp banking account is prohibited",
		})
	}

	err := companyService.DeleteCompany(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return ctx.NoContent(http.StatusNoContent)
}

func UpdateCompany(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if userID != id && role == "company" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other comp banking account is prohibited",
		})
	}

	updatedCompany := new(company.Company)

	if err := ctx.Bind(updatedCompany); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	comp, err := companyService.GetCompanyByID(userID)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	company.UpdateCompanyInfo(comp, updatedCompany)

	if !company.IsValid(updatedCompany) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid company data",
		})
	}

	err = companyService.UpdateCompany(updatedCompany)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updatedCompany)
}
