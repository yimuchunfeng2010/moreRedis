package global

import (
	"github.com/garyburd/redigo/redis"
	"more-for-redis/types"
	pb "more-for-redis/more_rpc/more_proto"
	"google.golang.org/grpc"
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
}{
	ZkIPaddr:         "127.0.0.1:2181",
	RedisAddr:        "127.0.0.1:6379",
	LocalRpcAddr:     "127.0.0.1:50051",
	RemoteRpcServers: []string{"127.0.0.1:50051"},
	RedisConn:        nil,
	Timeout:          5,
	PreDoReqList:     nil,
	RpcConn:          nil,
	RpcClient:        nil,
}
