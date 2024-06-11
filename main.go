package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"log"
	postgres "main/src/config/database/postgres/gorm"
	"main/src/controller/routes"
)

func init() {

}

func main() {

	var tools toolkit.Tools
	var logger toolkit.Logger
	config, err := tools.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read the env file")
	}

	database, err := postgres.NewPostgresGORMConnection(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("Error ao conectar no no banco, error=%s", err.Error())
		return
	} else {
		fmt.Println("conexao com sucesso")
	}

	userController := initDependencies(database)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	logger.Info("Iniciando")
}
