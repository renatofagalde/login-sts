package model

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"time"
)

const TOKEN_TYPE = "Bearer"

var (
	TOKEN_INVALIDO = "token inválido"
)

func (ud *userDomain) GenerateToken() (string, *toolkit.RestErr) {
	var tools toolkit.Tools
	config, err := tools.LoadConfig(".")

	var rest_err toolkit.RestErr

	//quais campos dentro do jwt
	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"exp":   time.Now().Add(time.Hour * 8).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.TokenSymmetricKey))
	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("erro ao gerar jwt token, err=%s", err.Error()))
	}
	return fmt.Sprintf("%s %s", TOKEN_TYPE, tokenString), nil
}
