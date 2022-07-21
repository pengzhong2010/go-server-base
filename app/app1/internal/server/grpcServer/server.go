package grpcServer

import (
	"context"
	"log"
	"net"

	pb "github.com/pengzhong2010/go-server-base/app/app1/api/v1"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetRobot(ctx context.Context, in *pb.IDReq) (*pb.Robot, error) {
	return &pb.Robot{}, nil
}
func (s *server) CreateRobot(context.Context, *pb.RobotCReq) (*pb.UpdateResp, error) {
	return &pb.UpdateResp{}, nil
}
func (s *server) UpdateRobot(context.Context, *pb.RobotUReq) (*pb.UpdateResp, error) {
	return &pb.UpdateResp{}, nil
}
func (s *server) DeleteRobot(context.Context, *pb.IDReq) (*pb.UpdateResp, error) {
	return &pb.UpdateResp{}, nil
}
func (s *server) ListRobot(context.Context, *pb.ListRobotReq) (*pb.ListRobotResp, error) {
	return &pb.ListRobotResp{}, nil
}

func New() (s *grpc.Server, cf func(), err error) {
	lis, errT := net.Listen("tcp", port)
	if errT != nil {
		err = errT
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s = grpc.NewServer()
	pb.RegisterRobotAPIServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return
}
