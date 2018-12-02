package utils

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

// Config is viper`s instance
var Config *viper.Viper

// Logger is logrus`s instance
var Logger *logrus.Logger

func init() {
	runMode := os.Getenv("env")
	var configName string
	switch runMode {
	case "stage":
		configName = "statge"
	case "test":
		configName = "test"
	case "prod":
		configName = "prod"
	default:
		configName = "dev"
	}
	initConfig("./conf", configName)
	initLog(runMode)
}

func initConfig(path, name string) {
	Config = viper.New()
	Config.AddConfigPath(path)
	Config.SetConfigName(name)
	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initLog(env string) {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})
	switch env {
	case "dev":
		Logger.SetOutput(os.Stdout)
	default:
		Logger.SetOutput(os.Stdout)
	}
}
