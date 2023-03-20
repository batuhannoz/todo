package store

import (
	"fmt"
	"github.com/batuhannoz/todo/backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(config *config.MySQL) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
