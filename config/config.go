package config

import (
	"os"
	"sync"

	"github.com/jinzhu/configor"
)

var (
	conf *Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "development"
		}

		conf = &Config{}
		configor.Load(conf, "config/config.yml")
		conf.ENV = env
	})
	return conf
}
