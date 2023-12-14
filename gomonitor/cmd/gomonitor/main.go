package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/kedare/gomonitor/pb"
	"github.com/mackerelio/go-osstat/cpu"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type GrpcServer struct {
	pb.UnimplementedMonitoringServiceServer
}

// GetCpuUsageInfo returns cpu usage info via grPC (one shot)
func (s *GrpcServer) GetCpuUsageInfo(ctx context.Context, in *pb.CpuUsageInfoRequest) (*pb.CpuUsageInfoResponse, error) {
	log.Println("GetCpuUsageInfo")
	return GetCpuUsageInfo(in.Interval)
}

// StreamCpuUsageInfo streams cpu usage info via grPC
func (s *GrpcServer) StreamCpuUsageInfo(in *pb.CpuUsageInfoRequest, srv pb.MonitoringService_StreamCpuUsageInfoServer) error {
	log.Println("StreamCpuUsageInfo")
	for {
		err := srv.Context().Err()
		if err != nil {
			log.Println("End of streaming, reason: ", err.Error())
			return err
		}
		cpuUsageInfo, err := GetCpuUsageInfo(in.Interval)
		if err != nil {
			return err
		}
		if err := srv.Send(cpuUsageInfo); err != nil {
			return err
		}
		log.Println("Sent cpu usage info in stream")
	}
}

// GetCpuUsageInfo returns cpu usage info
func GetCpuUsageInfo(interval int32) (*pb.CpuUsageInfoResponse, error) {
	before, err := cpu.Get()
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Duration(interval) * time.Second)
	after, err := cpu.Get()

	total := float64(after.Total - before.Total)
	cpuUser := float64(after.User-before.User) / total * 100
	cpuSystem := float64(after.System-before.System) / total * 100
	cpuIdle := float64(after.Idle-before.Idle) / total * 100

	return &pb.CpuUsageInfoResponse{
		SystemTime: int32(cpuSystem),
		UserTime:   int32(cpuUser),
		IdleTime:   int32(cpuIdle),
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
