package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	Admin    AdminConfig    `yaml:"admin"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

type DatabaseConfig struct {
	Type   string        `yaml:"type"`
	Sqlite SqliteConfig  `yaml:"sqlite"`
}

type SqliteConfig struct {
	Path string `yaml:"path"`
}

type JWTConfig struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

type AdminConfig struct {
	DefaultUsername string `yaml:"default_username"`
	DefaultPassword string `yaml:"default_password"`
}

var AppConfig *Config

func InitConfig() error {
	data, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		return err
	}

	AppConfig = &Config{}
	err = yaml.Unmarshal(data, AppConfig)
	if err != nil {
		return err
	}

	return nil
}
