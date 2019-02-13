package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"moreRedis/data"
	"net/http"
	"moreRedis/global"
	"time"
)

func Set(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("Set Key:%s, Value:%s\n", key, value)
	//获取本地写锁
	startTime := time.Now()
	global.Config.LocalRWLocker.Lock()
	defer func(){
		global.Config.LocalRWLocker.Unlock()
		durationTime := time.Now().Sub(startTime)
		logrus.Infof("Get durationTime %d",durationTime)
	}()
	err := data.Set(key,value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "success"})
	}
}
