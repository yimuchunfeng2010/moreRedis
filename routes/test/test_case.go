package test

import (
	//"moreRedis/routes/rpc"
	"moreRedis/global"
	"github.com/garyburd/redigo/redis"
	pb "moreRedis/more_rpc/more_proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"moreRedis/distributed_lock"
	"fmt"
	"sync"
	"moreRedis/more_rpc"
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

	// 初始化本地读写锁
	global.Config.LocalRWLocker = new(sync.RWMutex)
	if nil == global.Config.LocalRWLocker {
		logrus.Errorf("init LocalRWLocker Fail")
		return
	}

	go more_rpc.MoreRpcInit()

}
func MoreSetGet() (err error) {
	// 测试读写比例10:1
	//readCnt := 10
	//writeCnt := 1
	//key := "AAAAAAA"
	//value := "BBBBBB"
	//go func(){
	//	for i := 0; i < writeCnt; i++{
	//rpc.Set(key, value)
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
	//key := "AAAAAAA"
	//value := "BBBBBB"
	//go func(){
	//for i := 0; i < writeCnt; i++{
	//redis_operation.RedisSet(key, value)
	//}
	//}()
	//go func(){
	//for i := 0; i < readCnt; i++{
	//redis_operation.RedisGet(key)
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
	//var testCnt int64 = 100000
	//var i int64
	//var rpcSetTime int64
	//var rpcGetTime int64
	//
	//var RedisSetTime int64
	//var RedisGetTime int64
	//
	//var LightetTime int64
	//// 测试redis直接读写性能
	//
	//for i = 0; i < testCnt; i++ {
	//	startTime := time.Now()
	//	err := rpc.RedisSet("AAAA", "BBBB")
	//	if err != nil {
	//		logrus.Errorf("rpc.Set err:%s", err.Error())
	//	}
	//	durTime := time.Now().Sub(startTime)
	//	intTime := int64(durTime)
	//	RedisSetTime += intTime
	//}
	//
	//for i = 0; i < testCnt; i++ {
	//	startTime := time.Now()
	//	_,  err := rpc.RedisGet("AAAA")
	//	if err != nil {
	//		logrus.Errorf("rpc.Set err:%s", err.Error())
	//	}
	//	durTime := time.Now().Sub(startTime)
	//	intTime := int64(durTime)
	//	RedisGetTime += intTime
	//}
	//
	//
	//// 测试轻量级读性能
	//for i = 0; i < testCnt; i++ {
	//	startTime := time.Now()
	//	_, err := rpc.LigthGet("AAAA")
	//	if err != nil {
	//		logrus.Errorf("rpc.Set err:%s", err.Error())
	//	}
	//	durTime := time.Now().Sub(startTime)
	//	intTime := int64(durTime)
	//	LightetTime += intTime
	//}
	//
	//// 测试本方案读写性能
	//for i = 0; i < testCnt; i++ {
	//	startTime := time.Now()
	//	_, err := rpc.Get("AAAA")
	//	if err != nil {
	//		logrus.Errorf("rpc.Set err:%s", err.Error())
	//	}
	//	durTime := time.Now().Sub(startTime)
	//	intTime := int64(durTime)
	//	rpcGetTime += intTime
	//}
	//
	//for i = 0; i < testCnt/10; i++ {
	//	startTime := time.Now()
	//	err := rpc.Set("AAAA", "BBBB")
	//	if err != nil {
	//		logrus.Errorf("rpc.Set err:%s", err.Error())
	//	}
	//	durTime := time.Now().Sub(startTime)
	//	intTime := int64(durTime)
	//	rpcSetTime += intTime
	//}
	//
	//logrus.Infof("Redis Set Time %dns", RedisSetTime/testCnt)
	//logrus.Infof("Redis Get Time %dns", RedisGetTime/testCnt)
	//logrus.Infof("My Set Time %dns", rpcSetTime/testCnt)
	//logrus.Infof("My Get Time %dns", rpcGetTime/testCnt)
	//logrus.Infof("My LigthGet Time %dns", LightetTime/testCnt)
}
