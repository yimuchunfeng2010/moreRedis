package more_for_redis

import ("github.com/gin-gonic/gin"
"more-for-redis/routes/rest")

func init(){

}


func main(){
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