package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

var (
	Param = AppConfig()
)

type RedisConfig struct {
	Address      string `toml:"address"`
	AuthPassword string `toml:"auth_password"`
}

type Config struct {
	Redis RedisConfig
}

func AppConfig() Config {
	var config Config
	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		log.Println(err)
	}

	return config
}
