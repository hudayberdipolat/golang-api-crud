package dbConfig

import (
	"fmt"
	"github.com/hudayberdipolat/golang-api-crud/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbConfig(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DbConfig.DbHost,
		cfg.DbConfig.DbUser,
		cfg.DbConfig.DbPassword,
		cfg.DbConfig.DbName,
		cfg.DbConfig.DbPort,
		cfg.DbConfig.DbSllMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
