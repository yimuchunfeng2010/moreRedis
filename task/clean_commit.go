package task

import (
	"github.com/sirupsen/logrus"
	"more-for-redis/services"
	"more-for-redis/global"
	"fmt"
	"time"
	"more-for-redis/types"
)

// 清理过期commit
func CleanCommit() {
	spec := global.Config.CleanCommitCronSpec
	err := services.AddCrontab(spec, DoCleanCommit)
	if err != nil {
		logrus.Info(fmt.Sprintf("Start Cron spec[%s] name[%s] failed", spec, "DoCleanCommit"))
	} else {
		logrus.Info(fmt.Sprintf("Start Cron spec[%s] name[%s] success", spec, "DoCleanCommit"))
	}
}

func DoCleanCommit() {
	if nil == global.Config.PreDoReqList {
		return
	}

	for (time.Now().Unix()-global.Config.PreDoReqList.CreateTime >= global.Config.CommitTimeout) {
		global.Config.PreDoReqList = global.Config.PreDoReqList.Next
	}

	if nil == global.Config.PreDoReqList.Next{
		return
	}
	var preNode *types.ProcessingRequest = global.Config.PreDoReqList
	var curNode *types.ProcessingRequest = global.Config.PreDoReqList.Next
	for nil != curNode {
		if time.Now().Unix()-curNode.CreateTime < global.Config.CommitTimeout {
			preNode = curNode
			curNode = curNode.Next
		} else {
			preNode.Next = curNode.Next
			curNode = curNode.Next

		}

	}

	return
}
