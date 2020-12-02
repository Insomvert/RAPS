package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func BuildDbConfig() *DbConfig {
	dbConfig := DbConfig{
		Host:     "localhost", // mysql host
		Port:     3306,        // mysql port
		User:     "root",      // mysql username
		Password: "",          // mysql password
		DbName:   "raps",      // database name
	}
	return &dbConfig
}

func DbURL(dbConfig *DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)
}
