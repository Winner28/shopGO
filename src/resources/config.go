package resources

import (
	"encoding/json"
	"log"
	"os"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Dialect  string `json:"dialect"`
}

func GetDBConfig() *DBConfig {
	file, err := os.Open("./resources/config/dbProperties.json")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var config DBConfig
	if err = decoder.Decode(&config); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	return &config
}
