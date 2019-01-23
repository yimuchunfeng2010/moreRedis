package internal_interface

import (
	"more-for-redis/types"
	"more-for-redis/global"
	"more-for-redis/redis_operation"
	"errors"
	"time"
)

func PreSet(key string, value string, commitID int64)(err error){
	todoReq := &types.ProcessingRequest{
		CommitID:commitID,
		Req:types.ReqType_SET,
		Key:key,
		Value:value,
		Next:nil,
		CreateTime:time.Now().Unix(),
	}

	// 插入到头部
	todoReq.Next = global.Config.PreDoReqList

	global.Config.PreDoReqList = todoReq

	return
}


func Commit(commitID int64)(err error){
	if nil == global.Config.PreDoReqList{
		return errors.New("commit job not exist")
	}

	var curJob  *types.ProcessingRequest = nil
	if(commitID == global.Config.PreDoReqList.CommitID){
		curJob = global.Config.PreDoReqList
		global.Config.PreDoReqList = global.Config.PreDoReqList.Next
	} else{
		tmpPre := global.Config.PreDoReqList
		for tmpPre.Next != nil && tmpPre.Next.CommitID != commitID{
			tmpPre = tmpPre.Next
		}
		if tmpPre.Next != nil && tmpPre.Next.CommitID == commitID{
			curJob = tmpPre.Next
			tmpPre.Next = tmpPre.Next.Next
		}

	}

	if nil == curJob{
		return errors.New("commit job not exist")
	} else{
		return DoJob(curJob)
	}
	return
}

func Drop(commitID int64)(err error){
	if(commitID == global.Config.PreDoReqList.CommitID){
		global.Config.PreDoReqList = global.Config.PreDoReqList.Next
	} else{
		tmpPre := global.Config.PreDoReqList
		for tmpPre.Next != nil && tmpPre.Next.CommitID != commitID{
			tmpPre = tmpPre.Next
		}
		if tmpPre.Next != nil && tmpPre.Next.CommitID == commitID{
			tmpPre.Next = tmpPre.Next.Next
		}

	}
	return
}

func DoJob(job *types.ProcessingRequest)(err error){
	if nil == job {
		return errors.New("job nil")
	}

	switch job.Req {
	case types.ReqType_SET:
		err = redis_operation.RedisSet(job.Key,job.Value)
	default:
		err = errors.New("Wrong Request Type")
	}
	return
}