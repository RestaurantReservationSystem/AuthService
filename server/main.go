package main

import (
	"auth_service/config"
	pb "auth_service/genproto"
	"auth_service/service"
	"auth_service/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatal("error->%s", err.Error())
	}
	cnf := config.Load()
	listen, err := net.Listen("tcp", cnf.HTTPPort)
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, service.NewUserService(postgres.NewUserRepository(db)))
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("error->%s", err.Error())
	}
}
