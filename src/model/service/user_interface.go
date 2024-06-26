package service

import (
	toolkit "github.com/renatofagalde/golang-toolkit"
	"main/src/model"
	"main/src/model/repository"
)

func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{repository: repository}
}

type userDomainService struct {
	repository repository.UserRepository
}

type UserDomainService interface {
	FindUserByEmail(email string) (model.UserDomainInterface, *toolkit.RestErr)
	LoginService(domainInterface model.UserDomainInterface) (model.UserDomainInterface, string, *toolkit.RestErr)
	TokenVerify(string) *toolkit.RestErr
}
