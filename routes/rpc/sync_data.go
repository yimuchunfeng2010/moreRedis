package rpc
import (
	"more-for-redis/data"
	"more-for-redis/global"
	"github.com/sirupsen/logrus"
)

// 从主节点同步数据
func SyncData()( err error) {
	global.Config.LocalRWLocker.Lock()
	defer func() {
		global.Config.LocalRWLocker.Unlock()
	}()

	// 从主节点同步数据
	err = data.SyncData()
	if err != nil {
		logrus.Errorf("SyncData failed [err:%s]",err.Error())
	}
	return

}
