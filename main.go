package more_for_redis

import (
	"github.com/gin-gonic/gin"
	"more-for-redis/routes/rest"
	"more-for-redis/global"
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
	"more-for-redis/more_rpc"
)

func init() {
	conn, err := redis.Dial("tcp", global.Config.RedisAddr)
	if err != nil {
		logrus.Errorf("connect redis error %s",err.Error())
		return
	}
	global.Config.RedisConn = conn
}

func main() {

	go more_rpc.MoreRpcInit()
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/:key", rest.Get)
		v1.DELETE("/:key", rest.Delete)
		v1.POST("/:key/:value", rest.Set)
		v1.PUT("/:key/:value", rest.Update)

	}

	router.Run(":8000")
}
