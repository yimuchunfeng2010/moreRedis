package main

import (
	"more-for-redis/routes/rpc"
	"more-for-redis/redis_operation"
	"more-for-redis/global"
	"github.com/garyburd/redigo/redis"
	pb "more-for-redis/more_rpc/more_proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"more-for-redis/distributed_lock"
	"fmt"
	"sync"
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
	var hosts = []string{global.Config.ZkIPaddr}
	global.Config.ZkConn, _, err = zk.Connect(hosts, 100000*time.Minute)
	if err != nil {
		logrus.Errorf("Connect %s", err.Error())
		return
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

func ZkConnTest() (err error) {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lockName, err := distributed_lock.Lock()
			if err != nil {
				logrus.Errorf(fmt.Sprintf("services.Lock Failed! [Err:%s]", err.Error()))
				return

			}
			logrus.Infof("lockName %s", lockName)
			err = distributed_lock.Unlock(lockName)
			if err != nil {
				logrus.Errorf(fmt.Sprintf("services.UnLock Failed! [Err:%s]", err.Error()))
				return

			}
		}()
	}
	wg.Wait()
	return
}

func main() {
	err := ZkConnTest()
	if err != nil {
		logrus.Errorf("Test Fail Err: %s", err.Error())
	} else {
		logrus.Infof("Test Success")
	}
}
