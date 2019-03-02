package config

import (
	"log"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/viper"
)

var config *viper.Viper

func ReadConfigFile(env string) {
	config = viper.New()

	pwd, err := os.Getwd()
	if err != nil {
		glog.Fatalf("Error get current path, %s", err)
	}

	config.SetConfigFile(pwd + "/config/" + env + ".json")
	config.AddConfigPath(pwd + "/config")

	config.SetConfigName(env)
	config.SetConfigType("json")
	// Searches for config file in given paths and read it
	if err := config.ReadInConfig(); err != nil {
		glog.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	log.Println("Using config: ", config.ConfigFileUsed())
	log.Println("name", config.GetString("name"))

}

func GetConfig() *viper.Viper {
	return config
}
