package internal_interface

import (
	"moreRedis/types"
	"moreRedis/global"
	"moreRedis/internal_interface/util"
	"moreRedis/redis_operation"
	"errors"
	"time"
	"github.com/sirupsen/logrus"
)

func PreSet(xKey string, xValue string, commitID int64) (err error) {
	// 数据解密
	var deKey []byte
	var deValue []byte
	if "" != global.Config.Aeskey{
		deKey, err = util.AesDecrypt([]byte(xKey), []byte(global.Config.Aeskey))
		if err != nil {
			logrus.Warningf("AesEncrypt.Key Failed! [Err:%s]", err.Error())
			return
		}

		deValue, err = util.AesDecrypt([]byte(xValue), []byte(global.Config.Aeskey))
		if err != nil {
			logrus.Warningf("AesEncrypt.value Failed! [Err:%s]", err.Error())
			return
		}
	} else{
		deKey = []byte(xKey)
		deValue = []byte(xValue)
	}

	todoReq := &types.ProcessingRequest{
		CommitID:   commitID,
		Req:        types.ReqType_SET,
		Key:        string(deKey),
		Value:      string(deValue),
		Next:       nil,
		CreateTime: time.Now().Unix(),
	}

	// 插入到头部
	todoReq.Next = global.Config.PreDoReqList

	global.Config.PreDoReqList = todoReq

	return
}

func Commit(commitID int64) (err error) {
	if nil == global.Config.PreDoReqList {
		return errors.New("commit job not exist")
	}

	var curJob *types.ProcessingRequest = nil
	if (commitID == global.Config.PreDoReqList.CommitID) {
		curJob = global.Config.PreDoReqList
		global.Config.PreDoReqList = global.Config.PreDoReqList.Next
	} else {
		tmpPre := global.Config.PreDoReqList
		for tmpPre.Next != nil && tmpPre.Next.CommitID != commitID {
			tmpPre = tmpPre.Next
		}
		if tmpPre.Next != nil && tmpPre.Next.CommitID == commitID {
			curJob = tmpPre.Next
			tmpPre.Next = tmpPre.Next.Next
		}

	}

	if nil == curJob {
		return errors.New("commit job not exist")
	} else {
		return DoJob(curJob)
	}
	return
}

func Drop(commitID int64) (err error) {
	if (commitID == global.Config.PreDoReqList.CommitID) {
		global.Config.PreDoReqList = global.Config.PreDoReqList.Next
	} else {
		tmpPre := global.Config.PreDoReqList
		for tmpPre.Next != nil && tmpPre.Next.CommitID != commitID {
			tmpPre = tmpPre.Next
		}
		if tmpPre.Next != nil && tmpPre.Next.CommitID == commitID {
			tmpPre.Next = tmpPre.Next.Next
		}

	}
	return
}

func DoJob(job *types.ProcessingRequest) (err error) {
	if nil == job {
		return errors.New("job nil")
	}

	switch job.Req {
	case types.ReqType_SET:
		err = redis_operation.RedisSet(job.Key, job.Value)
	default:
		err = errors.New("Wrong Request Type")
	}
	return
}
