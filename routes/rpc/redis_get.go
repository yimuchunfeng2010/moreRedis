package rpc
import (
	"more-for-redis/redis_operation"
)

func RedisGet(key string)(value string, err error) {

	// 直接读取redis数据
	value, err = redis_operation.RedisGet(key)

	return
}
