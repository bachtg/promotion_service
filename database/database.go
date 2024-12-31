package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"promotion_service/config"
	"promotion_service/internal/repository"
)

func InitDb(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&repository.Event{}, &repository.Voucher{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
