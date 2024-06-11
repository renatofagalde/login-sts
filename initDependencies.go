package main

import (
	"gorm.io/gorm"
	"main/src/controller/user"
	"main/src/model/repository"
	"main/src/model/service"
)

func initDependencies(database *gorm.DB) controller.UserControllerInterface {
	r := repository.NewUerRepository(database)
	domainService := service.NewUserDomainService(r)
	return controller.NewControllerInterface(domainService)
}
