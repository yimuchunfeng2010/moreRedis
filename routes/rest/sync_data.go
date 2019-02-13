package rest

import (
	"github.com/gin-gonic/gin"
	"moreRedis/data"
	"moreRedis/global"
	"net/http"
)

func SyncData(context *gin.Context) {

	//获取本地写锁
	global.Config.LocalRWLocker.Lock()
	defer func(){
		global.Config.LocalRWLocker.Unlock()
	}()
	err := data.SyncData()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "success"})
	}
}