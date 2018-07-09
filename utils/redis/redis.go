package redis

import (
	"fmt"
	"sync"

	gredis "github.com/go-redis/redis"
	"github.com/pinem/server/config"
)

var (
	once   sync.Once
	client *gredis.Client
)

func Get() *gredis.Client {
	once.Do(func() {
		conf := config.Get()
		client = gredis.NewClient(&gredis.Options{
			Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
			Password: conf.Redis.Password,
			DB:       conf.Redis.Db,
		})
	})
	return client
}
