package service

import (
	toolkit "github.com/renatofagalde/golang-toolkit"
	"main/src/model"
)

func (userDomainService *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *toolkit.RestErr) {
	return userDomainService.repository.FindUserByEmail(email)
}
