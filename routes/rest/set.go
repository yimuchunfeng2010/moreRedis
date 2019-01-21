package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"more-for-redis/data"
	"net/http"
)

func Set(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("Set Key:%s, Value:%s\n", key, value)
	//
	err := data.Set(key,value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "success"})
	}
}
