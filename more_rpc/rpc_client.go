package more_rpc

import (
	"time"
	pb "more-for-redis/more_rpc/more_proto"
	"context"
	"github.com/sirupsen/logrus"
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

