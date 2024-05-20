package main

import (
	"context"
	"log"
	"net"

	pb "helloworld_grpc/protos/helloworld_grpc"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port) // создаем прослушивание сети tcp на порту
	if err != nil {
		log.Fatalf("failed to listen: %v", err) // если ошибка прослушивания не nil - то падаем
	}
	s := grpc.NewServer()                         // создаем новый grpc сервер
	pb.RegisterGreeterServer(s, &GreeterServer{}) // регистрируем наш GreeterServer в новом grpc сервере
	log.Printf("Server listening on %s", port)    // логгируем что сервер начал работать (прослушивать) на нашем порту
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err) // если при запуске сервера появилась ошибка - выводи ее
	}
}
