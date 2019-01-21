package more_rpc

import (
	"net"
	"more-for-redis/global"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "more-for-redis/more_rpc/more_proto"
	"context"
)
type MoreServer struct {
}

func (s *MoreServer) InGetKey(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	return
}

func (s *MoreServer) InSetValue(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	return
}

func (s *MoreServer) Commit(ctx context.Context, data *pb.CommitIDMsg)(resp *pb.CommitIDMsg, err error){
	return
}

func (s *MoreServer) Drop(ctx context.Context, data *pb.CommitIDMsg)(resp *pb.CommitIDMsg, err error){
	return
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