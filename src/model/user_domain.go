package model

import (
	toolkit "github.com/renatofagalde/golang-toolkit"
	"golang.org/x/crypto/bcrypt"
)

type userDomain struct {
	id       string
	email    string
	password string
	name     string
}

func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}
func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) HashPassword() (string, *toolkit.RestErr) {
	var restErr toolkit.RestErr
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		return "", restErr.NewRestErr("Validation hash", "error validation", 500, nil)
	}

	return string(hashPassword), nil
}

func (ud *userDomain) CheckPassword(hashed string, password string) *toolkit.RestErr {
	var restErr toolkit.RestErr
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return restErr.NewRestErr("Validation hash", "error validation", 500, nil)
	}
	return nil
}
