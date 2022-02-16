package config

import (
	"os"
	"strconv"
	"sync"
)

//AppConfig Application configuration
type AppConfig struct {
	Port     int `yaml:"port"`
	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

//GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port, _ = strconv.Atoi(os.Getenv("Port"))
	defaultConfig.Database.Driver = os.Getenv("DB_Driver")
	defaultConfig.Database.Name = os.Getenv("DB_Name")
	defaultConfig.Database.Address = os.Getenv("DB_Address") //172.17.0.1
	defaultConfig.Database.Port, _ = strconv.Atoi(os.Getenv("DB_Port"))
	defaultConfig.Database.Username = os.Getenv("DB_Username")
	defaultConfig.Database.Password = os.Getenv("DB_Password")
	return &defaultConfig
}
