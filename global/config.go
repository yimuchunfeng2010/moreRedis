package global

import "github.com/garyburd/redigo/redis"

var Config = struct {
	ZkIPaddr     string
	RedisConn    redis.Conn
	RedisAddr    string
	LocalRpcAddr string
}{
	ZkIPaddr:     "192.168.228.143:2181",
	RedisAddr:    "192.168.228.143:6379",
	LocalRpcAddr: "127.0.0.1:50051",
	RedisConn:    nil,
}
