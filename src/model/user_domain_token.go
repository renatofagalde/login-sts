package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"os"
	"strings"
	"time"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
	TOKEN_INVALIDO = "token inválido"
)

func (ud *userDomain) GenerateToken() (string, *toolkit.RestErr) {

	//todo read from app.env
	secretKey := os.Getenv(JWT_SECRET_KEY)
	var rest_err toolkit.RestErr

	//quais campos dentro do jwt
	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"exp":   time.Now().Add(time.Hour * 8).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("erro ao gerar jwt token, err=%s", err.Error()))
	}
	return tokenString, nil
}

func VerifyTokenMiddleware(c *gin.Context) {
	secretKey := os.Getenv(JWT_SECRET_KEY)
	var rest_err toolkit.RestErr
	var logger toolkit.Logger

	tokenValue := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(RemoveBearer(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secretKey), nil
		}

		errRest := rest_err.NewUnauthorizedRequestError(TOKEN_INVALIDO)
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return nil, errRest
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedRequestError(TOKEN_INVALIDO)
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedRequestError(TOKEN_INVALIDO)
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
	}

	logger.Info(fmt.Sprintf("User authenticated %#v", userDomain))
}

func RemoveBearer(token string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(token, prefix) {
		token = strings.TrimPrefix(prefix, token)
	}
	return token
}
