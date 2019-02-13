package rest
import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"moreRedis/redis_operation"
	"net/http"
	"time"
)

func RedisSet(context *gin.Context) {

	key := context.Param("key")
	value := context.Param("value")
	startTime := time.Now()
	defer func(){
		durationTime := time.Now().Sub(startTime)
		logrus.Infof("Redis Set durationTime %+va",durationTime)
	}()
	logrus.Infof("Set Key:%s",key)
	// 直接设置redis数据
	err := redis_operation.RedisSet(key,value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "success","value":value})
	}
}
