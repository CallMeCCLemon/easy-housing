package service

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"latentlab.cc/easyhousing/api"
	"latentlab.cc/easyhousing/pkg/logging"
	"log"
	"net"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("server-port", 8889, "The server port")
)

type EasyHousingServerImpl struct {
	api.UnimplementedEasyHousingServer
}

func (e EasyHousingServerImpl) Echo(ctx context.Context, request *api.EchoRequest) (*api.EchoResponse, error) {
	logger := logging.CreateLogger()
	logger.Info(fmt.Sprintf("Echo Request: %s", request))
	return &api.EchoResponse{Message: request.Message}, nil
}

func (e EasyHousingServerImpl) RegisterUser(ctx context.Context, request *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (e EasyHousingServerImpl) AuthenticateUser(ctx context.Context, request *api.AuthenticateUserRequest) (*api.AuthenticateUserResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (e EasyHousingServerImpl) RegisterHome(ctx context.Context, request *api.RegisterHomeRequest) (*api.RegisterHomeResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (e EasyHousingServerImpl) AddDoc(ctx context.Context, request *api.AddDocRequest) (*api.AddDocResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (e EasyHousingServerImpl) ScheduleAppointment(ctx context.Context, request *api.ScheduleAppointmentRequest) (*api.ScheduleAppointmentResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (e EasyHousingServerImpl) CompleteListing(ctx context.Context, request *api.CompleteListingRequest) (*api.CompleteListingResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (e EasyHousingServerImpl) GetListingStatus(ctx context.Context, request *api.GetListingStatusRequest) (*api.GetListingStatusResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (e EasyHousingServerImpl) mustEmbedUnimplementedEasyHousingServer() {
	//TODO implement me
	panic("implement me")
}

func NewEasyHousingServer() *EasyHousingServerImpl {
	s := &EasyHousingServerImpl{}
	return s
}

func Start() {
	flag.Parse()
	log.Println(fmt.Sprintf("Starting Listening on port %d", *port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	api.RegisterEasyHousingServer(grpcServer, NewEasyHousingServer())
	grpc_health_v1.RegisterHealthServer(grpcServer, NewHealthServer())
	reflection.Register(grpcServer)
	log.Println("Serving gRPC traffic now")
	grpcServer.Serve(lis)
}

func NewHealthServer() *health.Server {
	s := health.NewServer()
	return s
}
