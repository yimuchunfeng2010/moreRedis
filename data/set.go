package data

import (
	"more-for-redis/distributed_lock"
	"more-for-redis/internal_interface/util"
	pb "more-for-redis/more_rpc/more_proto"
	"github.com/sirupsen/logrus"
	"more-for-redis/more_rpc"
	"time"
	"more-for-redis/global"
)

func Set(key string, value string)(err error){
	lockName, err := distributed_lock.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	defer func(){
		distributed_lock.Unlock(lockName)
	}()

	logrus.Infof("%s Set Key:%s, Value:%s\n", key, value)

	// 发起提议
	commitID, err := util.GenCommitID("Set" + key + value)
	if err != nil {
		return
	}
	ackChan := make(chan bool)
	for _, client := range  global.Config.RpcClient{
		go func(){
			_, err = more_rpc.SetValue(client,pb.Data{Key:key,Value:value,CommitID:commitID})
			if err != nil {
				ackChan <- true
			} else{
				ackChan <- false
			}
		}()
	}

	timeout := global.Config.Timeout
	ackCount := 0
	for timeout != 0 && ackCount < len(global.Config.RemoteRpcServers) {

		select {
		case _, ok := <-ackChan:
			if ok {
				ackCount++
			} else {
				goto ProcessJob
			}
		default:

		}

		time.Sleep(time.Second / 1000)
		timeout--
	}
	close(ackChan)

	ProcessJob:
	// 提交
	if ackCount == len(global.Config.RemoteRpcServers) {
		for _, client := range global.Config.RpcClient {
			go more_rpc.Commit(client, pb.CommitIDMsg{CommitID:commitID})
		}
	} else { //撤销任务
		for _, client := range global.Config.RpcClient {
			go more_rpc.Drop(client, pb.CommitIDMsg{CommitID:commitID})
		}
	}
	return
}