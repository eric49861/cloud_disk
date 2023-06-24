package config

import (
	"cloud-disk/rpc/user/model"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// InitMySQL 初始化MySQL连接
func InitMySQL(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败: ", err.Error())
		return nil
	} else {
		log.Println("初始化MySQL连接成功")
	}
	//创建表结构
	db.AutoMigrate(&model.User{})
	return db
}

// InitRedis 初始化Redis连接
// Redis用来缓存邮箱验证码
func InitRedis(addr, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("连接redis出错: ", err.Error(), "\n")
	} else {
		log.Println("初始化Redis连接成功")
	}
	return rdb
}

type Config struct {
	zrpc.RpcServerConf
	MySQLConfig struct {
		DSN string `yaml:"dsn" json:"dsn"`
	} `json:"MysqlConfig" yaml:"MysqlConfig"`
	RedisConfig struct {
		Addr     string `yaml:"addr" json:"addr"`
		Password string `yaml:"password" json:"password"`
		DB       int    `yaml:"db" json:"db"`
	} `json:"RedisConfig" yaml:"RedisConfig"`
}
