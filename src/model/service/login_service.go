package service

import (
	"fmt"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
	"main/src/model"
)

func (ud *userDomainService) LoginService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *toolkit.RestErr) {
	var logger toolkit.Logger

	logger.Info(fmt.Sprintf("Login: %s LoginService", userDomain.GetEmail()), zap.String("journey", "Login"))

	user, err := ud.FindUserByEmail(userDomain.GetEmail())
	if err != nil {
		//logger.Error("init loginUser erro ao validar", err)
		return nil, "", err
	}

	fmt.Println(user)

	err = userDomain.CheckPassword(user.GetPassword(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := userDomain.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	return user, token, nil

}
