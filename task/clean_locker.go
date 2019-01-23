package task

import (
	"github.com/sirupsen/logrus"
	"more-for-redis/services"
	"more-for-redis/global"
	"more-for-redis/distributed_lock"
	"fmt"
	"time"
	"strconv"
)

// 清理过期Locker
func CleanLocker() {
	spec := global.Config.CleanLockerCronSpec
	err := services.AddCrontab(spec, DoCleanLocker)
	if err != nil {
		logrus.Info(fmt.Sprintf("Start Cron spec[%s] name[%s] failed", spec, "DoCleanLocker"))
	} else {
		logrus.Info(fmt.Sprintf("Start Cron spec[%s] name[%s] success", spec, "DoCleanLocker"))
	}
}

func DoCleanLocker() {

	var dir string = "/Lock"
	// 获取当前子节点
	children, _, err := global.Config.ZkConn.Children(dir)
	if err != nil {
		logrus.Errorf("RLock Children %s", err.Error())
		return
	}

	for _, child := range children{
		nodeName := dir + "/" + child
		data, _, err := global.Config.ZkConn.Get(nodeName)
		if err != nil {
			logrus.Errorf("ZkConn.Get faild nodeName:%s, err:%s", nodeName, err.Error())
			continue
		}

		createTimeUnix, err := strconv.ParseInt(string(data), 10, 64)
		if err != nil {
			logrus.Errorf("strconv.ParseIntt faild data:%s, err:%s", data, err.Error())
			continue
		}


		// 超时则删除节点
		if time.Now().Unix() - createTimeUnix >= global.Config.LockerTimeout{
			distributed_lock.DeleteNode(nodeName)
		}


	}

	return
}