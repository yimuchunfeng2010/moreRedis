package rest
import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"moreRedis/redis_operation"
	"net/http"
	"time"
)

func RedisGet(context *gin.Context) {

	key := context.Param("key")
	logrus.Infof("Get Key:%s",key)
	startTime := time.Now()
	defer func(){
		durationTime := time.Now().Sub(startTime)
		logrus.Infof("Redis Get durationTime %+v",durationTime)
	}()
	// 直接读取redis数据
	value, err := redis_operation.RedisGet(key)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "success","value":value})
	}
}
