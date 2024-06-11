package service

import (
	"fmt"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
	"main/src/model"
)

func (userDomainService *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *toolkit.RestErr) {
	var logger toolkit.Logger

	logger.Info(fmt.Sprintf("FindUserByEmail: %s userDomainService", email), zap.String("journey", "FindUserByEmail"))

	user, err := userDomainService.repository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil

}
