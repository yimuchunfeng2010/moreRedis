package more_rpc

import (
	"net"
	"moreRedis/global"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "moreRedis/more_rpc/more_proto"
	"context"
	"moreRedis/internal_interface"
	"moreRedis/redis_operation"
	"time"
)

type MoreServer struct {
}

func (s *MoreServer) InGetKey(ctx context.Context, data *pb.Data) (resp *pb.Data, err error) {
	return
}

func (s *MoreServer) InSetValue(ctx context.Context, data *pb.Data) (resp *pb.Data, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer func() {
		cancel()
	}()

	err = internal_interface.PreSet(data.Key, data.Value, data.CommitID)
	if err != nil {
		logrus.Warningf("PreSet Failed[data:%+v, err:%+s]", data, err.Error())
		return
	}

	logrus.Infof("internal_interface.PreSet Success")
	return &pb.Data{}, nil
}

func (s *MoreServer) Commit(ctx context.Context, data *pb.CommitIDMsg) (resp *pb.CommitIDMsg, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer func() {
		cancel()
	}()
	logrus.Infof("internal_interface.Commit")
	err = internal_interface.Commit(data.CommitID)
	if err != nil {
		logrus.Warningf("Commit Failed[data:%+v, err:%+s]", data, err.Error())
		return
	}
	return &pb.CommitIDMsg{}, nil
}

func (s *MoreServer) Drop(ctx context.Context, data *pb.CommitIDMsg) (resp *pb.CommitIDMsg, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer func() {
		cancel()
	}()

	err = internal_interface.Drop(data.CommitID)
	if err != nil {
		logrus.Warningf("Drop Failed[data:%+v, err:%+s]", data, err.Error())
		return
	}
	return &pb.CommitIDMsg{}, nil
}

func (s *MoreServer) InGetKeys(data *pb.Data, stream pb.MoreRpcProto_InGetKeysServer) (err error) {
	keys, err := redis_operation.RedisGetKeys()
	if err != nil {
		logrus.Warningf("RedisGetKeys Failed[Err:%s]", err.Error())
		return err
	}
	for _, key := range keys {
		value, _ := redis_operation.RedisGet(key)
		if err = stream.Send(&pb.Data{Key: key, Value: value}); err != nil {
			logrus.Warnf("stream.Send Failed[key: %s,value: %s, Err:%s]", key, value, err.Error())
			return
		}
	}
	return
}

func MoreRpcInit() {
	lis, err := net.Listen("tcp", global.Config.LocalRpcAddr)
	if err != nil {
		logrus.Fatal("failed to listen: %s", err.Error())
		return
	}

	s := grpc.NewServer()
	pb.RegisterMoreRpcProtoServer(s, &MoreServer{})

	s.Serve(lis)
}
