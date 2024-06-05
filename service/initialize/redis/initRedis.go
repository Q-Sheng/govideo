package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"service/global"
)

func InitRedis(RedisCon *global.RedisConfig) *redis.Client {

	// todo 是否有必要采用go-retry来多次重新连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", RedisCon.DbHost, RedisCon.Port),
		Password: RedisCon.Password,
		DB:       RedisCon.Dbnum,
	})
	// check connection
	//context.Background() 创建了一个空的 Context
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()

	if err != nil {
		// fatal 返回错误信息并终止
		log.Fatalf("Failed to connect to Redis: %v", err)
		return nil

	} else {
		//// 测试操作rdb
		//// 读出
		//
		//val, err := rdb.Get(ctx, "mykey").Result()
		//if err == redis.Nil {
		//	fmt.Println("key2 does not exist")
		//} else if err != nil {
		//	panic(err)
		//} else {
		//	fmt.Println("mykey", val)
		//}

		return rdb
	}

	// 设置 并读出

}
