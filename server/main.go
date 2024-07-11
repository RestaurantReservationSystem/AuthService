package main

import (
	"auth_service/config"
	pb "auth_service/genproto"
	"auth_service/service"
	postgres "auth_service/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatal("error connection for database ")
	}
	cnf := config.Load()
	listen, err := net.Listen("tcp", cnf.HTTPPort)
	if err != nil {
		log.Fatalf("error-> tcp connection error->%s", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, service.NewUserService(postgres.NewUserRepository(db)))
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("error-> api_get_way connection error->%s", err.Error())
	}
}
