package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.BindEnv("config_use_env")
	useEnv := viper.GetBool("config_use_env")
	if useEnv {
		log.Println("loading config from env variables")
		viper.BindEnv("postgres_user")
		viper.BindEnv("postgres_password")
		viper.BindEnv("postgres_db")
		viper.BindEnv("postgres_hostname")
		viper.BindEnv("postgres_port")
		viper.BindEnv("postgres_ssl")

		viper.BindEnv("api_port")
		viper.BindEnv("api_audit")
	} else {
		log.Println("loading config from file config.yaml")
		viper.SetConfigName("config") // name of config file (without extension)
		viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(".")      // optionally look for config in the working directory
		err := viper.ReadInConfig()   // Find and read the config file
		if err != nil {               // Handle errors reading the config file
			log.Fatalf("failed to parse config file - %v", err)
		}
	}
}
