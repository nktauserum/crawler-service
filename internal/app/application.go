package app

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nktauserum/crawler-service/internal/handlers"
	"github.com/nktauserum/crawler-service/proto"
	"github.com/nktauserum/crawler-service/proto/pb"
	"google.golang.org/grpc"
)

type application struct {
	port int
	grpc *RPCServer
}

type RPCServer struct {
	server *proto.Server
	port   int
}

func NewRPCServer(port int) *RPCServer {
	return &RPCServer{
		port: port,
	}
}

func (s *RPCServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	s.server = &proto.Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, s.server)

	log.Printf("Starting gRPC server on port %d", s.port)
	return grpcServer.Serve(lis)
}

func NewApplication(port int) *application {
	return &application{port: port, grpc: NewRPCServer(50000)}
}

func (app *application) Run() error {
	go func() {
		if err := app.grpc.Start(); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	router := gin.Default()

	// Simple health check
	router.GET("/", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })

	authorized := router.Group("/")
	authorized.Use(handlers.CheckAPIToken())
	{
		authorized.POST("/crawl", handlers.Crawl)
		authorized.GET("/result", handlers.Result)
	}

	return router.Run(fmt.Sprintf(":%d", app.port))
}
