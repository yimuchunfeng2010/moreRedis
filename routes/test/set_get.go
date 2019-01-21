package test

import (
	"more-for-redis/routes/rpc"
	"more-for-redis/redis_operation"
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
func MoreSetGet() (err error) {
	// 测试读写比例10:1
	//readCnt := 10
	//writeCnt := 1
	key := "AAAAAAA"
	value := "BBBBBB"
	//go func(){
	//	for i := 0; i < writeCnt; i++{
	rpc.Set(key, value)
	//	}
	//}()
	//go func(){
	//	for i := 0; i < readCnt; i++{
	//rpc.Get(key)
	//	}
	//}()
	return
}

func RedisSetGet() (err error) {
	// 测试读写比例10:1
	//readCnt := 10
	//writeCnt := 1
	key := "AAAAAAA"
	value := "BBBBBB"
	//go func(){
	//for i := 0; i < writeCnt; i++{
	redis_operation.RedisSet(key, value)
	//}
	//}()
	//go func(){
	//for i := 0; i < readCnt; i++{
	redis_operation.RedisGet(key)
	//}
	//}()
	return
}

func main()  {
	err := MoreSetGet()
	if err != nil {
		logrus.Errorf("Test Fail Err: %s",err.Error())
	} else {
		logrus.Infof("Test MoreSetGet Success")
	}
}