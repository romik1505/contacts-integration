package server

import (
	"context"
	"flag"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
	"week3_docker/internal/handler"
	contact_service "week3_docker/internal/service/contact"
	contact "week3_docker/pkg/api/contact_service"
)

var (
	grpcPort = flag.String("grpc_port", ":50051", "grpc port")
	httpPort = flag.String("port", ":8080", "http port")
)

type Server struct {
	cs      contact_service.IService
	handler *handler.Handler
}

func NewServer(cs contact_service.IService, h *handler.Handler) Server {
	return Server{
		cs:      cs,
		handler: h,
	}
}

func (s Server) Run() {
	ctx := context.Background()
	go RunGRPCServer(ctx, s.cs)

	mux := RunGateway(ctx)
	if err := mux.HandlePath(http.MethodPost, "/api/contacts/sync", s.handler.ContactSync); err != nil {
		log.Fatalf("HandlePath /api/contacts/sync: %v", err)
	}

	if err := mux.HandlePath(http.MethodPost, "/api/account/{id}/contacts/hook", s.handler.ContactActionsHook); err != nil {
		log.Fatalf("HandlePath /api/account/{id}/contacts/hook%v", err)
	}
	mx := SwaggerMux(mux)

	withCors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(mx)

	if err := http.ListenAndServe(*httpPort, withCors); err != nil {
		log.Fatalf("error start server: %v", err)
	}
}

func SwaggerMux(mx *runtime.ServeMux) *http.ServeMux {
	httpMux := http.NewServeMux()

	httpMux.Handle("/", mx)
	httpMux.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/docs/docs.swagger.json"),
	))

	fileServer := http.FileServer(http.Dir("docs"))
	httpMux.Handle("/docs/", http.StripPrefix("/docs/", fileServer))
	return httpMux
}

func RunGRPCServer(ctx context.Context, cs contact_service.IService) {
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
		grpc_middleware.WithUnaryServerChain( //nolint
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)
	contact.RegisterContactServiceServer(gs, cs)
	if err = gs.Serve(lis); err != nil {
		log.Fatalf("RunGRPCServer Serve: %v", err)
		return
	}
}

func RunGateway(ctx context.Context) *runtime.ServeMux {
	m := &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}}

	mx := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, m),
	)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := contact.RegisterContactServiceHandlerFromEndpoint(ctx, mx, *grpcPort, opts); err != nil {
		log.Fatalf("RunGateway: RegisterContactServiceHandlerFromEndpoint: %v", err)
	}
	return mx
}
