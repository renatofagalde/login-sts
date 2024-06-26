package model

import (
	"github.com/goccy/go-json"
	toolkit "github.com/renatofagalde/golang-toolkit"
)

func (ud *userDomain) ToJSON() (string, *toolkit.RestErr) {
	var restErr toolkit.RestErr
	b, err := json.Marshal(ud)
	if err != nil {
		return "", restErr.NewInternalServerError("impossible convert to json")
	}
	return string(b), nil
}
