package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

var (
	config = PlogConfig()
)

type RedisConfig struct {
	Address      string `toml:"address"`
	AuthPassword string `toml:"auth_password"`
}

type Config struct {
	Redis RedisConfig
}

func PlogConfig() Config {
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Println(err)
	}

	return config
}
