package server

import (
	"context"
	"flag"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"week3_docker/internal/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
	contact "week3_docker/pkg/api/contact_service"
)

var (
	grpcPort = flag.String("grpc_port", ":50051", "grpc port")
	httpPort = flag.String("port", ":8080", "http port")
)

type Server struct {
	cs *service.ContactService
}

func NewServer(cs *service.ContactService) Server {
	return Server{
		cs: cs,
	}
}

func (s Server) Run() {
	ctx := context.Background()
	go RunGRPCServer(ctx, s.cs)

	mux := RunGateway(ctx)
	mx := SwaggerMux(ctx)
	mx.Handle("/", mux)

	if err := http.ListenAndServe(*httpPort, mx); err != nil {
		log.Fatal("error start server: %v", err)
	}
}

func SwaggerMux(ctx context.Context) *http.ServeMux {
	mx := http.NewServeMux()

	mx.HandleFunc("/swagger/docs.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/docs.swagger.json")
	})
	mx.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("docs.swagger.json"),
	))
	return mx
}

func RunGRPCServer(ctx context.Context, cs *service.ContactService) {
	lis, err := net.Listen("tcp", *grpcPort)
	if err != nil {
		log.Fatalf("error start grpc server %v", err)
	}

	customFunc := func(p interface{}) (err error) {
		log.Printf("panic triggered: %v", p)
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}
	gs := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)
	contact.RegisterContactServiceServer(gs, cs)
	gs.Serve(lis)
}

func RunGateway(ctx context.Context) *runtime.ServeMux {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}))

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	contact.RegisterContactServiceHandlerFromEndpoint(ctx, mux, *grpcPort, opts)
	return mux
}
