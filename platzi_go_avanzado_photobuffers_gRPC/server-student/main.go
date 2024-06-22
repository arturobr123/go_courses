package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"platzi.com/go/grpc/database"
	"platzi.com/go/grpc/server"
	studentpb "platzi.com/go/grpc/studentpb"
)

func main() {
	lis, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	server := server.NewServer(repo)

	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	grpcServer := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// INICIAR BASE DE DATOS
// docker build . -t platzi-grpc-db
// docker run -p 54321:5432 platzi-grpc-db

// INICIAR SERVIDOR
// go run server-student/main.go
