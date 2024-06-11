package postgres

import (
	"context"
	"fmt"
	toolkit "github.com/renatofagalde/golang-toolkit"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func NewPostgresGORMConnection(ctx context.Context, databaseSource string) (*gorm.DB, error) {

	var logger toolkit.Logger
	//todo remove password from log
	logger.Info(fmt.Sprintf("connection string: %v", databaseSource),
		zap.String("init", "NewPostgresGORMConnection"))

	db, err := gorm.Open(postgres.Open(databaseSource), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
