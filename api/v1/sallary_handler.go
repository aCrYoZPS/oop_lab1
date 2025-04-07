package api_v1

import (
	"net/http"
	"oopLab1/config"
	salaryprojects "oopLab1/core/salaryProjects"
	"oopLab1/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var salaryProjectService = salaryprojects.NewSalaryProjectService(config.GetConfig().Database)

func CreateSalaryProject(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)
	salaryProject := new(salaryprojects.SalaryProject)

	if role != "company" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "only company can issue salary project",
		})
	}

	if err := ctx.Bind(salaryProject); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	_, err := custService.GetCustomerByID(salaryProject.WorkerID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	salaryProject.ID = uuid.NewString()
	salaryProject.CompanyID = userID

	if !salaryprojects.IsValid(salaryProject) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid staff data",
		})
	}

	err = salaryProjectService.CreateSalaryProject(salaryProject)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"id": salaryProject.ID,
	})
}

func GetSalaryProject(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	salaryProject, err := salaryProjectService.GetSalaryProjectByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if salaryProject.CompanyID != userID && utils.GetPrivelegeLevel(role) < utils.GetPrivelegeLevel("manager") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to salary is prohibited",
		})
	}

	return ctx.JSON(http.StatusOK, salaryProject)
}

func GetAllSalaryProjects(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if utils.GetPrivelegeLevel(role) < utils.GetPrivelegeLevel("manager") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to salary is prohibited",
		})
	}

	salaryProjects, err := salaryProjectService.GetAllSalaryProjects()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, salaryProjects)
}

func GetAllSalaryProjectsByCompany(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if utils.GetPrivelegeLevel(role) < utils.GetPrivelegeLevel("company") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to salary is prohibited",
		})
	}

	salaryProjects, err := salaryProjectService.GetAllSalaryProjectsByCompany(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, salaryProjects)
}

func DeleteSalaryProject(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	salaryProject, err := salaryProjectService.GetSalaryProjectByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if salaryProject.CompanyID != userID && utils.GetPrivelegeLevel(role) < utils.GetPrivelegeLevel("manager") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to salary is prohibited",
		})
	}

	err = salaryProjectService.DeleteSalaryProject(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func UpdateSalaryProject(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	salaryProject, err := salaryProjectService.GetSalaryProjectByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if salaryProject.CompanyID != userID && utils.GetPrivelegeLevel(role) < utils.GetPrivelegeLevel("manager") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to salary is prohibited",
		})
	}

	updatedSalaryProject := &salaryprojects.SalaryProject{}

	if err := ctx.Bind(updatedSalaryProject); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	salaryprojects.UpdateSalaryProjectInfo(salaryProject, updatedSalaryProject)

	if !salaryprojects.IsValid(updatedSalaryProject) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid salary data",
		})
	}

	err = salaryProjectService.UpdateSalaryProject(updatedSalaryProject)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updatedSalaryProject)
}
