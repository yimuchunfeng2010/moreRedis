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
	ZkIPaddr            string
	RedisConn           redis.Conn
	RedisAddr           string
	LocalRpcAddr        string
	RemoteRpcServers    []string
	MasterServer        string
	Timeout             int
	PreDoReqList        *types.ProcessingRequest
	RpcConn             []*grpc.ClientConn
	MasterRpcConn       *grpc.ClientConn
	RpcClient           []pb.MoreRpcProtoClient
	MasterRpcClient     pb.MoreRpcProtoClient
	ZkConn              *zk.Conn
	ZkConnMaxTime       int
	LocalRWLocker       *sync.RWMutex
	CleanLockerCronSpec string
	CleanCommitCronSpec string
	LockerTimeout       int64
	CommitTimeout       int64
	Aeskey              string
}{
	ZkIPaddr:            "192.168.228.143:2181",
	RedisAddr:           "127.0.0.1:6379",
	LocalRpcAddr:        "192.168.228.143:50052",
	RemoteRpcServers:    []string{"192.168.228.143:50052", "192.168.228.142:50052", "192.168.228.143:50051"},
	MasterServer:        "192.168.228.143:50052",
	RedisConn:           nil,
	Timeout:             5000,
	PreDoReqList:        nil,
	RpcConn:             nil,
	MasterRpcConn:       nil,
	RpcClient:           nil,
	MasterRpcClient:     nil,
	ZkConn:              nil,
	ZkConnMaxTime:       1000000,
	LocalRWLocker:       nil,
	LockerTimeout:       10 * 60, // zk分布式锁10分钟超时则清理
	CommitTimeout:       10 * 60,
	CleanLockerCronSpec: "0 */10 * * * *",
	CleanCommitCronSpec: "0 */10 * * * *",
	Aeskey:              "01234567899801234567899801234567899876",
}
