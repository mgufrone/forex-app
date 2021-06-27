package service

import (
	"context"
	"fmt"
	"github.com/mgufrone/forex/internal/shared/infrastructure/healthcheck"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func NewDB() (*gorm.DB, healthcheck.HealthCheck, error) {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local&allowNativePasswords=true&allowOldPasswords=true",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_DB"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
		QueryFields:                              true,
	})
	return db, func(ctx context.Context) error {
		db1, err := db.DB()
		if err != nil {
			return err
		}
		return db1.Ping()
	}, err
}
