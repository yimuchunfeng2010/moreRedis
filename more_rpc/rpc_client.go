package more_rpc

import (
	"time"
	pb "more-for-redis/more_rpc/more_proto"
	"context"
	"github.com/sirupsen/logrus"
	"more-for-redis/redis_operation"
	"io"
)

func SetValue(client pb.MoreRpcProtoClient, data pb.Data) (resp *pb.Data, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer func(){
		cancel()
	}()

	if resp, err = client.InSetValue(ctx, &data); err != nil {
		logrus.Warningf("SetValue Failed[data:%+v, err:%+v]", data, err)
		return
	}
	return
}

func Commit(client pb.MoreRpcProtoClient, data pb.CommitIDMsg) (resp *pb.CommitIDMsg, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer func(){
		cancel()
	}()

	if resp, err = client.Commit(ctx, &data); err != nil {
		logrus.Warningf("Commit Failed[data:%+v, err:%+v]", data, err)
		return
	}
	return
}

func Drop(client pb.MoreRpcProtoClient, data pb.CommitIDMsg) (resp *pb.CommitIDMsg, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer func(){
		cancel()
	}()

	if resp, err = client.Drop(ctx, &data); err != nil {
		logrus.Warningf("Drop Failed[data:%+v, err:%+v]", data, err)
		return
	}
	return
}

func GetKeysAndSetInLocal(client pb.MoreRpcProtoClient)(err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	stream, err := client.InGetKeys(ctx,&pb.Data{})
	if err != nil{
		logrus.Warningf("InGetKeys Failed[err:%+s]", err.Error())
		return
	}
	for {
		data ,err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			logrus.Warningf("stream.Recv Failed[err:%s]",err.Error())
		}

		err = redis_operation.RedisSet(data.Key,data.Value)
		if err != nil {
			logrus.Warningf("RedisSet[data: %+v, err:%s]",data, err.Error())
		}
	}
	return
}

