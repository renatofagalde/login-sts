package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) TokenVerify(c *gin.Context) {
	var logger toolkit.Logger

	tokenValue := c.Request.Header.Get("Authorization")
	logger.Info(fmt.Sprintf("TokenVerify: %s", tokenValue), zap.String("journey", "TokenVerify"))

	err := uc.service.TokenVerify(tokenValue)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.Status(http.StatusOK)
}
