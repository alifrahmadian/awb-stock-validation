package tools

import (
	"context"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(ctx context.Context, dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when NewGormDB, error: %s\n", err.Error())
		return nil
	}

	return db.WithContext(ctx)
}
