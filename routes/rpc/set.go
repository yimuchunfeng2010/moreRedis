package rpc

import (
	"more-for-redis/data"
	"more-for-redis/global"
)

func Set(key string, value string) (err error) {
	//获取本地写锁
	global.Config.LocalRWLocker.Lock()
	defer func() {
		global.Config.LocalRWLocker.Unlock()
	}()
	err = data.Set(key, value)

	return
}
