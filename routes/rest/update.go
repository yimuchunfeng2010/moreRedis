package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"errors"
)

func Update(context *gin.Context) {

	key := context.Param("key")
	value := context.Param("value")
	logrus.Infof("Set Key:%s, Value:%s\n",key, value)

	// TODO 待实现分布式读写锁
	// TODO 待实现读写数据
	err := errors.New("no err")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "success"})
	}
}
