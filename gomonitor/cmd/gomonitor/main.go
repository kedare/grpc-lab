package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/kedare/gomonitor/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type GrpcServer struct {
	proto.UnimplementedMonitoringServiceServer
}

func (s *GrpcServer) GetCpuUsageInfo(ctx context.Context, in *proto.CpuUsageInfoRequest) (*proto.CpuUsageInfoResponse, error) {
	log.Println("GetCpuUsageInfo")
	return &proto.CpuUsageInfoResponse{
		SystemTime: 1,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterMonitoringServiceServer(s, &GrpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
