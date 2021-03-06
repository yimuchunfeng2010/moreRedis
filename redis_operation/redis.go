package redis_operation

import (
	"github.com/garyburd/redigo/redis"
	"moreRedis/global"
	"github.com/sirupsen/logrus"
)


func RedisSet(key string, value string)(err error){
	if nil == global.Config.RedisConn{
		logrus.Errorf("RedisConn nil")
		return
	}
	_, err = global.Config.RedisConn.Do("SET", key, value)
	if err != nil {
		logrus.Errorf("redis set error %s",err.Error())
		return
	}
	return
}
func RedisGet(key string)(value string, err error){
	if nil == global.Config.RedisConn{
		logrus.Errorf("RedisConn nil")
		return
	}
	value, err = redis.String(global.Config.RedisConn.Do("Get", key))
	if err != nil {
		logrus.Errorf("redis Get error %s",err.Error())
		return
	}
	return
}

func RedisGetKeys()(keys []string,err error){
	if nil == global.Config.RedisConn{
		logrus.Errorf("RedisConn nil")
		return
	}
	keys, err = redis.Strings(global.Config.RedisConn.Do("KEYS", "*"))
	if err != nil {
		logrus.Errorf("redis Get Keys error %s",err.Error())
		return
	}
	return
}