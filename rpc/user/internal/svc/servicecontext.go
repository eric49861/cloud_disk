package svc

import (
	"cloud-disk/rpc/user/internal/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	SQL    *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	println("addr:", c.RedisConfig.Addr)
	return &ServiceContext{
		Config: c,
		SQL:    config.InitMySQL(c.MySQLConfig.DSN),
		RDB:    config.InitRedis(c.RedisConfig.Addr, c.RedisConfig.Password, c.RedisConfig.DB),
	}
}
