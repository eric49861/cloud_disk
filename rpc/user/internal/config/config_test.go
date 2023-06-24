package config

import (
	"context"
	"testing"
	"time"
)

func TestInitMySQL(t *testing.T) {
	const dsn = "eric:8520@tcp(127.0.0.1:3306)/cloud_disk?charset=utf8mb4&parseTime=True&loc=Local"
	db := InitMySQL(dsn)
	if db == nil {
		t.Fatal("数据库连接测试失败")
	}
}

func TestInitRedis(t *testing.T) {
	rdb := InitRedis("topc.fun:6379", "8520", 0)
	if rdb == nil {
		t.Fatal("Redis连接测试失败")
	} else {
		rdb.Set(context.Background(), "code", "0000", time.Second*30)
		code, err := rdb.Get(context.Background(), "code").Result()
		if err != nil {
			t.Fatal("Redis连接测试失败")
			return
		}
		t.Log(code)
		t.Log("Redis连接测试成功")
	}
}
