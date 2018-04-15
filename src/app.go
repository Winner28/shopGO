package main

import (
	"fmt"
	"model"
	"resources"

	"github.com/jinzhu/gorm"
	// psql
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Work")
	config := resources.GetDBConfig()
	user := model.User{}
	user.FirstName = "Sasha"
	user.LastName = " sashyla"
	user.Email = "@mail"
	user.PasswordHash = "12345"
	dataSource := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.User, config.Password, config.Database, config.Host, config.Port)

	db, err := gorm.Open(config.Dialect, dataSource)
	if err != nil {
		panic(err)
	}

	err = db.Save(&user).Error
	if err != nil {
		panic(err)
	}

}
