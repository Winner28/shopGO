package service

import (
	"fmt"
	"resources"

	// psql & gorm
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type connection struct {
	DB *gorm.DB
}

var c *connection

func initialize() {
	c = new(connection)
	config := resources.GetDBConfig()
	dataSource := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.User, config.Password, config.Database, config.Host, config.Port)

	db, err := gorm.Open(config.Dialect, dataSource)
	if err != nil {
		panic(err)
	}
	c.DB = db
}

func getConnection() *connection {
	if c == nil {
		initialize()
		return c
	}
	return c
}
