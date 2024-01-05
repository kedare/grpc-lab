package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/kedare/gomonitor/pb"
	"github.com/mackerelio/go-osstat/cpu"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

var (
	port   = flag.Int("port", 50051, "The server port")
	tracer = otel.Tracer("gomonitor/main")
)

type GrpcServer struct {
	pb.UnimplementedMonitoringServiceServer
}

// GetCpuUsageInfo returns cpu usage info via grPC (one shot)
func (s *GrpcServer) GetCpuUsageInfo(ctx context.Context, in *pb.CpuUsageInfoRequest) (*pb.CpuUsageInfoResponse, error) {
	log.Println("GetCpuUsageInfo")
	return GetCpuUsageInfo(ctx, in.Interval)
}

// StreamCpuUsageInfo streams cpu usage info via grPC
func (s *GrpcServer) StreamCpuUsageInfo(in *pb.CpuUsageInfoRequest, srv pb.MonitoringService_StreamCpuUsageInfoServer) error {
	log.Println("StreamCpuUsageInfo")
	ctx := srv.Context()
	for {
		cpuUsageInfo, err := GetCpuUsageInfo(ctx, in.Interval)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			log.Println("End of streaming, reason: ", ctx.Err().Error())
			return ctx.Err()

		default:
			if err := srv.Send(cpuUsageInfo); err != nil {
				return err
			}
			log.Println("Sent cpu usage info in stream")
		}
	}
}

// GetCpuUsageInfo returns cpu usage info
func GetCpuUsageInfo(ctx context.Context, interval int32) (*pb.CpuUsageInfoResponse, error) {
	_, span := tracer.Start(ctx, "GetCpuUsageInfo")
	defer span.End()
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

func setupOpenTelemetry(ctx context.Context) (*sdkTrace.TracerProvider, error) {
	appResource, err := resource.New(
		ctx,
		resource.WithFromEnv(),
		resource.WithContainer(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
	)

	if err != nil {
		log.WithError(err).Error("Failed to initialize OpenTelemetry resource")
		return nil, err
	}

	consoleExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())

	client := otlptracehttp.NewClient()
	exporter, err := otlptrace.New(ctx, client)

	if err != nil {
		log.WithError(err).Error("Failed to initialize OpenTelemetry export pipeline")
		return nil, err
	}

	tracerProvider := sdkTrace.NewTracerProvider(
		sdkTrace.WithSampler(sdkTrace.AlwaysSample()),
		sdkTrace.WithBatcher(exporter),
		sdkTrace.WithBatcher(consoleExporter),
		sdkTrace.WithResource(appResource),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tracerProvider, nil
}

func main() {
	ctx := context.Background()
	_, err := setupOpenTelemetry(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to initialize OpenTelemetry")
		return
	}
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	pb.RegisterMonitoringServiceServer(s, &GrpcServer{})
	log.Printf("server listening at %v", lis.Addr())

	_, span := tracer.Start(ctx, "testing")
	span.End()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
