package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
	"main/src/controller/model/request"
	"main/src/model"
	"main/src/view"
	"net/http"
)

func (uc *userControllerInterface) Login(c *gin.Context) {
	var logger toolkit.Logger
	var restErr toolkit.RestErr
	var userRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errorMessage := restErr.NewBadRequestError("Erro ao validar user")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)
	domainResult, token, err := uc.service.LoginService(domain)
	if err != nil {
		errorMessage := restErr.NewBadRequestError("Erro ao validar o login")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	logger.Info(fmt.Sprintf("Login: %s userControllerInterface OK", userRequest.Email), zap.String("journey", "Login"))
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
