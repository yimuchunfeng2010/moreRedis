package rpc

import (
	"github.com/sirupsen/logrus"
	"more-for-redis/data"
)

func Set(key string, value string) (err error) {
	logrus.Infof("%s Set Key:%s, Value:%s\n", key, value)
	err = data.Set(key, value)
	return
}
