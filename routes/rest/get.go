package rest
import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"more-for-redis/redis_operation"
	"net/http"
)

func Get(context *gin.Context) {

	key := context.Param("key")
	logrus.Infof("Get Key:%s",key)


	value, err := redis_operation.RedisGet(key)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "success","value":value})
	}
}
