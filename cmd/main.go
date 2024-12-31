package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "promotion_service/api"
	"promotion_service/config"
	"promotion_service/database"
	"promotion_service/internal/repository"
	"promotion_service/internal/service"
)

func main() {
	var (
		cfg         = config.LoadConfig()
		gormDb, err = database.InitDb(cfg)
	)
	if err != nil {
		panic(err)
	}

	var (
		gormRepo = repository.NewRepository(gormDb)
		service  = service.NewService(nil, gormRepo)
		address  = fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	)

	// running gRPC server
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("failed to listen", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterLoyaltyServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	go func() {
		httpMux := runtime.NewServeMux()

		err = api.RegisterLoyaltyServiceHandlerFromEndpoint(
			context.Background(), httpMux, address, []grpc.DialOption{grpc.WithInsecure()},
		)
		if err != nil {
			log.Fatal("failed to register HTTP gateway", err)
		}

		httpAddress := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.HTTPPort)

		log.Print("Starting HTTP server on: ", httpAddress)

		if err := http.ListenAndServe(httpAddress, httpMux); err != nil {
			log.Fatal("HTTP server failed", err)
		}
	}()

	log.Print("Server's running on: ", cfg.Server.Host+":"+cfg.Server.Port)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
}
