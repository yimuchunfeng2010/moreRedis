package global

import (
	"github.com/garyburd/redigo/redis"
	"more-for-redis/types"
	pb "more-for-redis/more_rpc/more_proto"
	"github.com/samuel/go-zookeeper/zk"
	"google.golang.org/grpc"
	"sync"
)

var Config = struct {
	ZkIPaddr         string
	RedisConn        redis.Conn
	RedisAddr        string
	LocalRpcAddr     string
	RemoteRpcServers []string
	Timeout          int
	PreDoReqList     *types.ProcessingRequest
	RpcConn          []*grpc.ClientConn
	RpcClient        []pb.MoreRpcProtoClient
	ZkConn           *zk.Conn
	ZkConnMaxTime    int
	LocalRWLocker    *sync.RWMutex
}{
	ZkIPaddr:         "192.168.228.143:2181",
	RedisAddr:        "127.0.0.1:6379",
	LocalRpcAddr:     "192.168.228.143:50052",
	RemoteRpcServers: []string{"192.168.228.143:50052"},
	RedisConn:        nil,
	Timeout:          5000,
	PreDoReqList:     nil,
	RpcConn:          nil,
	RpcClient:        nil,
	ZkConn:           nil,
	ZkConnMaxTime:    1000000,
	LocalRWLocker:    nil,
}
