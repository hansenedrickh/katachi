package dependencies

import (
	"fmt"

	"github.com/hansenedrickh/katachi/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		logrus.Fatalf("[Database] Failed to Open Database Connection: %v", err.Error())
		return nil, err
	}

	return db, nil
}
