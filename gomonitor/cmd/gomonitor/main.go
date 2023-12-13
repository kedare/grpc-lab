package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/kedare/gomonitor/pb"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type GrpcServer struct {
	pb.UnimplementedMonitoringServiceServer
}

func (s *GrpcServer) GetCpuUsageInfo(ctx context.Context, in *pb.CpuUsageInfoRequest) (*pb.CpuUsageInfoResponse, error) {
	log.Println("GetCpuUsageInfo")
	return GetCpuUsageInfo(in.Interval)
}

func (s *GrpcServer) StreamCpuUsageInfo(in *pb.CpuUsageInfoRequest, srv pb.MonitoringService_StreamCpuUsageInfoServer) error {
	log.Println("StreamCpuUsageInfo")
	for {
		cpuUsageInfo, err := GetCpuUsageInfo(in.Interval)
		if err != nil {
			return err
		}
		if err := srv.Send(cpuUsageInfo); err != nil {
			return err
		}
		log.Println("Sent cpu usage info in stream")
	}
	log.Println("End of streaming")
	return nil
}

func GetCpuUsageInfo(interval int32) (*pb.CpuUsageInfoResponse, error) {

	// Generate a random number between 0 and 100, until we have a real implementation
	cpuUsage := rand.Intn(100)

	time.Sleep(time.Duration(interval) * time.Second)

	return &pb.CpuUsageInfoResponse{
		SystemTime: int32(cpuUsage),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMonitoringServiceServer(s, &GrpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
