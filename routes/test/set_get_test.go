package test

import (
	"testing"
	"more-for-redis/global"
	"github.com/garyburd/redigo/redis"
	pb "more-for-redis/more_rpc/more_proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

}

func Benchmark_MoreSetGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MoreSetGet()
	}

}

func Benchmark_RedisSetGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RedisSetGet()
	}
}
