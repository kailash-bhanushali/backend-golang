package config

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)


type ServerConfig struct {
	Port     string `mapstructure:"PORT"`
	LogLevel string `mapstructure:"LOGLEVEL"`
	Env      string `mapstructure:"ENV"`
}

func NewServerConfig() *ServerConfig {

	env := gin.Mode()
	serverconfig := &ServerConfig{}
	var result map[string]interface{}

	if env == "release" || env == "debug" {
		viper.SetConfigType("env")
		viper.AutomaticEnv()
		envMap := map[string]string{
			"LogLevel": "LOGLEVEL",
			"Port": "PORT",
		}
		for path, env := range envMap{
			if err := viper.BindEnv(path, env); err != nil{
				log.WithError(err).WithField("env", env).Panic("Failed to Bind Env Var")
			}
		}
		if err := viper.ReadInConfig(); err != nil{
			log.Info("Failed to read Config File, will read from env vars")
		}
		if err := viper.Unmarshal(&result); err != nil{
			log.Panic("Failed to Unmarshal data")
		}
		decErr := mapstructure.Decode(result, serverconfig)
		if decErr != nil{
			log.Panic("Error Decoding")
		}
		log.Println("Printing Server Config Details: ", serverconfig)
		return serverconfig
	} else {
		log.Panic("Error loading Config")
	}
	return serverconfig
}
