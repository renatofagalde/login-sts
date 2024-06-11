package controller

import (
	"github.com/gin-gonic/gin"
	"main/src/model/service"
)

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
}

func NewControllerInterface(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{service}
}

type userControllerInterface struct {
	service service.UserDomainService
}
