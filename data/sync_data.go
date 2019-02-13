package data

import (
	"moreRedis/distributed_lock"
	"github.com/sirupsen/logrus"
	"moreRedis/more_rpc"
	"moreRedis/global"
)

func SyncData()(err error){
	lockName, err := distributed_lock.Lock()
	if err != nil {
		logrus.Warningf("services.Lock Failed! [Err:%s]", err.Error())
		return

	}
	logrus.Infof("lockName %s", lockName)
	defer func() {
		distributed_lock.Unlock(lockName)
	}()

	if nil != global.Config.MasterRpcClient{
		err = more_rpc.GetKeysAndSetInLocal(global.Config.MasterRpcClient)
		if err != nil {
			logrus.Warningf("GetKeysAndSetInLocal Failed! [Err:%s]", err.Error())
			return

		}
	}

	return
}
