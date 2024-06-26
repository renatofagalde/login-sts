package service

import (
	"github.com/golang-jwt/jwt"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"strings"
)

var (
	TOKEN_INVALIDO = "token inválido"
)

func (ud *userDomainService) TokenVerify(tokenValue string) *toolkit.RestErr {

	var rest_err toolkit.RestErr
	var tools toolkit.Tools
	config, err := tools.LoadConfig(".")

	token, err := jwt.Parse(RemoveBearer(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(config.TokenSymmetricKey), nil
		}
		errRest := rest_err.NewUnauthorizedRequestError(TOKEN_INVALIDO)
		return nil, errRest
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedRequestError(TOKEN_INVALIDO)
		return errRest
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedRequestError(TOKEN_INVALIDO)
		return errRest
	}

	return nil

}
func RemoveBearer(token string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(token, prefix) {
		token = strings.TrimPrefix(token, prefix)
	}
	return token
}
