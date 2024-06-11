package model

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
	GetPassword() string
	SetID(string)
	//GenerateToken() (string, *toolkit.RestErr)
}
