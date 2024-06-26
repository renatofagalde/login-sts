package model

import toolkit "github.com/renatofagalde/golang-toolkit"

func NewUserDomain(id string, email string, password string, name string) UserDomainInterface {
	return &userDomain{
		id:       id,
		email:    email,
		password: password,
		name:     name,
	}
}

func NewUserLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password}
}

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetName() string
	GetPassword() string
	SetID(string)
	ToJSON() (string, *toolkit.RestErr)

	HashPassword() (string, *toolkit.RestErr)
	CheckPassword(string, string) *toolkit.RestErr

	GenerateToken() (string, *toolkit.RestErr)
}
