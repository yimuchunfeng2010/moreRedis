package rpc

import (
	"moreRedis/redis_operation"
	"moreRedis/global"
)

func Get(key string) (value string,  err error) {
	// 获取被动读锁
	global.Config.LocalRWLocker.RLock()
	defer func(){
		global.Config.LocalRWLocker.RUnlock()
	}()
	value, err = redis_operation.RedisGet(key)

	return
}
