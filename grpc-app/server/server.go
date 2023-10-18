package main

import (
	"context"
	"fmt"
	proto "grpc-app/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type appServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

// implementation based on proto.AppServiceServer interface
func (asi *appServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Add req received for x = %d and y = %d\n", x, y)
	result := x + y

	// create the response object
	addResponse := &proto.AddResponse{
		Result: result,
	}

	fmt.Printf("Sending add response [result = %d]\n", result)
	return addResponse, nil
}

func main() {
	asi := &appServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
