package more_rpc

import (
	"net"
	"more-for-redis/global"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "more-for-redis/more_rpc/more_proto"
	"context"
	"more-for-redis/internal_interface"
	"time"
)
type MoreServer struct {
}

func (s *MoreServer) InGetKey(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	return
}

func (s *MoreServer) InSetValue(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer func(){
		cancel()
	}()

	err = internal_interface.PreSet(data.Key,data.Value, data.CommitID)
	if err != nil {
		logrus.Warningf("PreSet Failed[data:%+v, err:%+s]", data, err.Error())
		return
	}

	logrus.Infof("internal_interface.PreSet Success")
	return &pb.Data{},nil
}

func (s *MoreServer) Commit(ctx context.Context, data *pb.CommitIDMsg)(resp *pb.CommitIDMsg, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer func(){
		cancel()
	}()
	logrus.Infof("internal_interface.Commit")
	err = internal_interface.Commit(data.CommitID)
	if err != nil {
		logrus.Warningf("Commit Failed[data:%+v, err:%+s]", data, err.Error())
		return
	}
	return &pb.CommitIDMsg{},nil
}

func (s *MoreServer) Drop(ctx context.Context, data *pb.CommitIDMsg)(resp *pb.CommitIDMsg, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer func(){
		cancel()
	}()

	err = internal_interface.Drop(data.CommitID)
	if err != nil {
		logrus.Warningf("Drop Failed[data:%+v, err:%+s]", data, err.Error())
		return
	}
	return &pb.CommitIDMsg{},nil
}

func MoreRpcInit(){
	lis, err := net.Listen("tcp", global.Config.LocalRpcAddr)
	if err != nil {
		logrus.Fatal("failed to listen: %s", err.Error())
		return
	}

	s := grpc.NewServer()
	pb.RegisterMoreRpcProtoServer(s, &MoreServer{})

	s.Serve(lis)
}