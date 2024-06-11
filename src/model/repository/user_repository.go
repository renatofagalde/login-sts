package repository

import (
	toolkit "github.com/renatofagalde/golang-toolkit"
	"gorm.io/gorm"
	"main/src/model"
)

type UserRepository interface {
	FindUserByEmail(email string) (model.UserDomainInterface, *toolkit.RestErr)
}

type userRepository struct {
	database *gorm.DB
}

func NewUerRepository(database *gorm.DB) UserRepository {
	return &userRepository{database}
}
