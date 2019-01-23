package services

import (
"github.com/robfig/cron"
"git.chinawayltd.com/golib/tools/timelib"
"errors"
)

var (
	// 执行定时调度
	jobCrontabPtr *cron.Cron
	// 记录
	uniqueCrontabID uint64
)

// 初始化Cron
func InitCrontab() error {
	if jobCrontabPtr == nil {
		jobCrontabPtr = cron.NewWithLocation(timelib.CST)
	}
	return nil
}

// 运行Cron
func RunCrontab() error {
	if jobCrontabPtr == nil {
		return errors.New("jobCrontabPtr is nil")
	}
	jobCrontabPtr.Start()

	return nil
}

// 增加作业至周期调度
func AddCrontab(spec string, cmd func()) error {
	if jobCrontabPtr == nil {
		return errors.New("jobCrontabPtr is nil")
	}
	return jobCrontabPtr.AddFunc(spec, cmd)
}

