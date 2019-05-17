package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type DatabaseConfig struct {
	user         string
	pwd          string
	url          string
	maxIdleConns int
	maxOpenConns int
}

func init() {
	var err error
	var configs = buildDatabaseConfig()
	db, err = gorm.Open("mysql", configs.user+":"+configs.pwd+"@/"+configs.url)
	if err != nil {
		panic(err)
	} else {
		db.DB().SetMaxIdleConns(configs.maxIdleConns)
		db.DB().SetMaxOpenConns(configs.maxOpenConns)
	}
}

func buildDatabaseConfig() DatabaseConfig {
	var config = new(DatabaseConfig)
	config.user = "root"
	config.pwd = "123123123"
	config.url = "localhost:3306/test?charset=utf8mb4&parseTime=true&loc=Local"
	config.maxOpenConns = 8
	config.maxIdleConns = 8
	return *config
}
