package main

import (
	"context"
	"log"

	pb "helloworld_grpc/protos/helloworld_grpc"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // устанавливаем связь с grpc сервером
	if err != nil {
		log.Fatalf("could not connect: %v", err) // обработка ошибки
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn) // создаем клиент gRPC

	name := "Mechislav"
	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name}) // реализация SayHello
	if err != nil {
		log.Fatalf("error calling SayHello: %v", err)
	}

	log.Printf("Response from server: %s", response.Message) // ответ сервера
}
