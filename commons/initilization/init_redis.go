package initilization

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
	"ksd-social-api/commons/global"
	"time"
)

// go get github.com/beego/beego/v2/client/rdb/redis@v2.0.6
// redigo官方文档：https://pkg.go.dev/github.com/gomodule/redigo/redis#pkg-examples
// go get github.com/gomodule/redigo/redis

// 创建缓存容器
var Rdb redis.Conn
var keyPrefix string

// 连接池连接
func PoolConnect() redis.Conn {
	redisHost := beego.AppConfig.DefaultString("redis.host", "127.0.0.1")
	redisPort := beego.AppConfig.DefaultInt("redis.port", 6379)
	dataBase := beego.AppConfig.DefaultInt("redis.database", 1)
	password := beego.AppConfig.DefaultString("redis.password", "")
	maxIdle := beego.AppConfig.DefaultInt("redis.maxIdle", 1)
	maxActive := beego.AppConfig.DefaultInt("redis.MaxActive", 10)
	timeout := beego.AppConfig.DefaultString("redis.timeout", "10")

	duration, _ := time.ParseDuration(timeout)
	pool := &redis.Pool{
		MaxIdle:     maxIdle,           // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   maxActive,         // 最大的激活连接数，表示同时最多有N个连接
		IdleTimeout: 300 * time.Second, // 最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Wait:        true,              // 超过最大连接数的操作:等待
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", redisHost, redisPort),
				redis.DialDatabase(dataBase),
				redis.DialPassword(password),
				redis.DialReadTimeout(duration*time.Second),
				redis.DialConnectTimeout(duration*time.Second),
			)
			if err != nil {
				logs.Error("缓存初始化失败....")
				return nil, err
			}
			//选择分区
			return c, nil
		},
	}

	return pool.Get()
}

// 初始化缓存
func InitRedis() {
	// 获取配置文件的前缀
	global.RedisKey = beego.AppConfig.DefaultString("redis.key", "ksd:")
}
