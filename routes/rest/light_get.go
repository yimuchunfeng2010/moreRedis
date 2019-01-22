package rest
import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"more-for-redis/redis_operation"
	"net/http"
	"time"
)

func LightGet(context *gin.Context) {

	key := context.Param("key")
	logrus.Infof("Get Key:%s",key)
	startTime := time.Now()
	defer func(){
		durationTime := time.Now().Sub(startTime)
		logrus.Infof("Light Get durationTime %+v",durationTime)
	}()
	// 不获取本地读锁，采用最终一致性方案
	value, err := redis_operation.RedisGet(key)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "success","value":value})
	}
}
