package main

import (
	"github.com/gin-gonic/gin"
	"moreRedis/routes/rest"
	"moreRedis/distributed_lock"
	"moreRedis/global"
	"github.com/garyburd/redigo/redis"
	pb "moreRedis/more_rpc/more_proto"
	"github.com/sirupsen/logrus"
	"moreRedis/more_rpc"
	"google.golang.org/grpc"
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"sync"
	"moreRedis/services"
	"moreRedis/task"
)

func init() {
	conn, err := redis.Dial("tcp", global.Config.RedisAddr)
	if err != nil {
		logrus.Errorf("connect redis error %s", err.Error())
		return
	}
	global.Config.RedisConn = conn

	for _, server := range global.Config.RemoteRpcServers {
		conn, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			logrus.Warningf("fail to dial: %s", err.Error())
			return
		}
		global.Config.RpcConn = append(global.Config.RpcConn, conn)
		global.Config.RpcClient = append(global.Config.RpcClient, pb.NewMoreRpcProtoClient(conn))
	}

	MasterConn, err := grpc.Dial(global.Config.MasterServer, grpc.WithInsecure())
	if err != nil {
		logrus.Warningf("fail to dial: %s", err.Error())
		return
	}
	global.Config.MasterRpcConn = MasterConn
	global.Config.MasterRpcClient = pb.NewMoreRpcProtoClient(MasterConn)

	var hosts = []string{global.Config.ZkIPaddr}
	global.Config.ZkConn, _, err = zk.Connect(hosts, 100000*time.Minute)
	if err != nil {
		logrus.Errorf("Connect %s", err.Error())
		return
	}

	// 初始化本地读写锁
	global.Config.LocalRWLocker = new(sync.RWMutex)
	if nil == global.Config.LocalRWLocker {
		logrus.Errorf("init LocalRWLocker Fail")
		return
	}

	// 初始化删除zookeeper锁目录下所有子节点
	err = distributed_lock.DeleteAllChildren("/Lock")
	if err != nil {
		logrus.Errorf("DeleteAllChildren error %s", err.Error())
		return
	}

}

func CronInit() {
	services.InitCrontab()
	services.RunCrontab()

	// 清理超时锁
	task.CleanLocker()
	// 清理超时Commit
	task.CleanCommit()
}

func main() {
	go more_rpc.MoreRpcInit()
	go CronInit()

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/:key", rest.Get)
		v1.POST("/:key/:value", rest.Set)
		//v1.POST("/syncData", rest.SyncData)

	}

	light := router.Group("/light")
	{
		light.GET("/:key", rest.LightGet)
	}

	redis := router.Group("/redis")
	{
		redis.GET("/:key", rest.RedisGet)
		redis.POST("/:key/:value", rest.RedisSet)
	}
	router.Run(":8000")
}
