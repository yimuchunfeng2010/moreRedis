package main

import (
	"github.com/gin-gonic/gin"
	"more-for-redis/routes/rest"
	"more-for-redis/global"
	"github.com/garyburd/redigo/redis"
	pb "more-for-redis/more_rpc/more_proto"
	"github.com/sirupsen/logrus"
	"more-for-redis/more_rpc"
	"google.golang.org/grpc"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func init() {
	conn, err := redis.Dial("tcp", global.Config.RedisAddr)
	if err != nil {
		logrus.Errorf("connect redis error %s",err.Error())
		return
	}
	global.Config.RedisConn = conn

	for _, server := range global.Config.RemoteRpcServers{
		conn, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			logrus.Warningf("fail to dial: %s", err.Error())
			return
		}
		global.Config.RpcConn = append(global.Config.RpcConn, conn)
		global.Config.RpcClient = append(global.Config.RpcClient,pb.NewMoreRpcProtoClient(conn))
	}

	var hosts = []string{global.Config.ZkIPaddr}
	global.Config.ZkConn, _, err = zk.Connect(hosts, 100000*time.Minute)
	if err != nil {
		logrus.Errorf("Connect %s", err.Error())
		return
	}

}

func main() {
	go more_rpc.MoreRpcInit()

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/:key", rest.Get)
		v1.DELETE("/:key", rest.Delete)
		v1.POST("/:key/:value", rest.Set)
		v1.PUT("/:key/:value", rest.Update)

	}

	router.Run(":8000")
}
