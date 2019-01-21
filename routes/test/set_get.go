package test

import (
	"more-for-redis/routes/rpc"
	"more-for-redis/redis_operation"
)

func MoreSetGet() (err error) {
    // 测试读写比例10:1
    //readCnt := 10
    //writeCnt := 1
	key := "AAAAAAA"
	value := "BBBBBB"
	//go func(){
	//	for i := 0; i < writeCnt; i++{
			rpc.Set(key, value)
	//	}
	//}()
	//go func(){
	//	for i := 0; i < readCnt; i++{
			rpc.Get(key)
	//	}
	//}()
	return
}


func RedisSetGet()(err error){
	// 测试读写比例10:1
	//readCnt := 10
	//writeCnt := 1
	key := "AAAAAAA"
	value := "BBBBBB"
	//go func(){
		//for i := 0; i < writeCnt; i++{
			redis_operation.RedisSet(key, value)
		//}
	//}()
	//go func(){
		//for i := 0; i < readCnt; i++{
			redis_operation.RedisGet(key)
		//}
	//}()
	return
}