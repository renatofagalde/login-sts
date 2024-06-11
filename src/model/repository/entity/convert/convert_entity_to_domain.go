package convert

import (
	"main/src/model"
	"main/src/model/repository/entity"
)

func ConvertEntityToDomain(userEntity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(userEntity.Email, userEntity.Password, userEntity.Name)
	return domain
}
