package model

import toolkit "github.com/renatofagalde/golang-toolkit"

func NewUserDomain(email string, password string, name string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
	}
}

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetName() string
	SetID(string)
	GenerateToken() (string, *toolkit.RestErr)
}
