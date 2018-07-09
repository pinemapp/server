package config

import (
	"fmt"
)

type Config struct {
	ENV    string
	App    string      `yaml:"app"`
	Host   string      `yaml:"host"`
	Port   int         `yaml:"port"`
	Secret string      `yaml:"secret"`
	DB     dbStruct    `yaml:"db"`
	I18n   i18nStruct  `yaml:"i18n"`
	Redis  redisStruct `yaml:"redis"`
	Token  tokenStruct `yaml:"token"`
}

type dbStruct struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:""password`
}

type i18nStruct struct {
	Default string `yaml:"default"`
	Dir     string `yaml:"dir"`
}

type redisStruct struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Db       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type tokenStruct struct {
	ExpiresAt int64 `yaml:"expires_at"`
}

func (c *Config) ServerAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) DbString() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		c.DB.Host, c.DB.Port, c.DB.Name, c.DB.User, c.DB.Password)
}
