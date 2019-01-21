package rpc

import (
	"github.com/sirupsen/logrus"
	"more-for-redis/redis_operation"
)

func Get(key string) (value string, err error) {
	// 直接从本地redis读取数据
	logrus.Infof("Get Key:%s\n", key)
	value, err = redis_operation.RedisGet(key)
	return
}
