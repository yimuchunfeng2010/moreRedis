package data

import (
	"moreRedis/distributed_lock"
	"moreRedis/internal_interface/util"
	pb "moreRedis/more_rpc/more_proto"
	"github.com/sirupsen/logrus"
	"moreRedis/more_rpc"
	"time"
	"moreRedis/global"
	"sync"
)

func Set(key string, value string) (err error) {
	lockName, err := distributed_lock.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	logrus.Infof("lockName %s", lockName)
	ackChan := make(chan bool, len(global.Config.RemoteRpcServers))
	defer func() {
		distributed_lock.Unlock(lockName)
		// TODO 处理通道关闭检测的问题
		close(ackChan)
	}()

	logrus.Infof("Set Key:%s, Value:%s\n", key, value)

	// 发起提议
	commitID, err := util.GenCommitID("Set" + key + value)
	if err != nil {
		return
	}

	var xKey []byte
	var xValue []byte
	if "" != global.Config.Aeskey {
		xKey, err = util.AesEncrypt([]byte(key), []byte(global.Config.Aeskey))
		if err != nil {
			logrus.Warningf("AesEncrypt.Key Failed! [Err:%s]", err.Error())
			return
		}

		xValue, err = util.AesEncrypt([]byte(value), []byte(global.Config.Aeskey))
		if err != nil {
			logrus.Warningf("AesEncrypt.Key Failed! [Err:%s]", err.Error())
			return

		}
	} else {
		xKey = []byte(key)
		xKey = []byte(value)
	}

	var wg sync.WaitGroup
	for _, client := range global.Config.RpcClient {
		go func() {
			wg.Add(1)
			_, err = more_rpc.SetValue(client, pb.Data{Key: string(xKey), Value: string(xValue), CommitID: commitID})
			if err != nil {
				ackChan <- true
			} else {
				ackChan <- false
			}
			wg.Done()
		}()
	}

	wg.Wait()
	timeout := global.Config.Timeout
	ackCount := 0
	for timeout != 0 && ackCount < len(global.Config.RemoteRpcServers) {

		select {
		case _, ok := <-ackChan:
			if ok {
				ackCount++
			}
		default:
		}

		time.Sleep(time.Millisecond)
		timeout--
	}

	// 提交
	if ackCount == len(global.Config.RemoteRpcServers) {
		for _, client := range global.Config.RpcClient {
			go more_rpc.Commit(client, pb.CommitIDMsg{CommitID: commitID})
		}
	} else { //撤销任务
		for _, client := range global.Config.RpcClient {
			go more_rpc.Drop(client, pb.CommitIDMsg{CommitID: commitID})
		}
	}
	return
}
