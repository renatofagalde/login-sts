package repository

import (
	"fmt"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
	"main/src/model"
	"main/src/model/repository/entity"
	"main/src/model/repository/entity/convert"
)

func (userRepository *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *toolkit.RestErr) {

	var logger toolkit.Logger
	var rest_err toolkit.RestErr
	var user entity.UserEntity
	//err := userRepository.database.Where("email =?", email).First(&user).Error
	//if err != nil {
	//	errorMessage := fmt.Sprintf("Site not found with this ID: %s", id)
	//	logger.Error(fmt.Sprintf("FindUserByEmail: %s userRepository  %+v", email.errorMessage), err, zap.String("journey", "FindUserByEmail"))
	//	return nil, rest_err.NewNotFoundError(errorMessage)
	//}

	query := fmt.Sprintf("select id_usuario,txt_email,txt_password,nom_usuario from usuarios_mst "+
		"where cod_usuario='%s' limit 1", email)
	logger.Info(fmt.Sprintf("FindUserByEmail: query %s", query),
		zap.String("journey", "FindUserByEmail"))
	rows, err := userRepository.database.Raw(query).Rows()
	if err != nil {
		errorMessage := fmt.Sprintf("Site not found with this email: %s", email)
		logger.Error(fmt.Sprintf("FindUserByEmail: %s userRepository  %+v", email, errorMessage),
			err, zap.String("journey", "FindUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name); err != nil {
			errorMessage := fmt.Sprintf("FindUserByEmail not found with this Email: %s", email)
			logger.Error(fmt.Sprintf("FindUserByEmail: %s userRepository  %+v", email, errorMessage),
				err, zap.String("journey", "FindUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)

		}
	}
	logger.Info(fmt.Sprintf("FindUserByEmail: %s userRepository  %+v", email, user),
		zap.String("journey", "FindUserByEmail"))

	return convert.ConvertEntityToDomain(user), nil
}
