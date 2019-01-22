package rpc
import (
	"more-for-redis/redis_operation"
)

func RedisSet(key string, value string)( err error) {

	// 直接设置redis数据
	err = redis_operation.RedisSet(key,value)

	return

}
