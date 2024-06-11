package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
	"main/src/view"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	var logger toolkit.Logger
	var rest_err toolkit.RestErr

	email := c.Param("email")
	logger.Info(fmt.Sprintf("FindUserByEmail: %s userControllerInterface", email), zap.String("journey", "FindUserByEmail"))
	if _, err := mail.ParseAddress(email); err != nil {
		message := "Email não é válido"
		errorMessage := rest_err.NewBadRequestError(message)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		message := "Erro ao recuperar email"
		logger.Error(message, err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logger.Info(fmt.Sprintf("FindUserByEmail: %s sucesso", email), zap.String("journey", "FindUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}
