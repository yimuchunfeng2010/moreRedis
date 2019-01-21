package redis_operation

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"more-for-redis/global"
	"github.com/sirupsen/logrus"
)

func main() {
	conn, err := redis.Dial("tcp", "10.1.210.69:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("SET", "name", "wd")
	if err != nil {
		fmt.Println("redis set error:", err)
	}
	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Got name: %s \n", name)
	}
}

func RedisSet(key string, value string)(err error){
	_, err = global.Config.RedisConn.Do("SET", key, value)
	if err != nil {
		logrus.Errorf("redis set error %s",err.Error())
		return
	}
	return
}
func RedisGet(key string)(value string, err error){
	value, err = redis.String(global.Config.RedisConn.Do("Get", key))
	if err != nil {
		logrus.Errorf("redis Get error %s",err.Error())
		return
	}
	return
}