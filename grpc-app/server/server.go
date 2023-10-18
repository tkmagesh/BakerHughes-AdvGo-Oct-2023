package main

import (
	"context"
	"fmt"
	proto "grpc-app/proto"
	"log"
	"net"
	"time"

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

func (asi *appServiceImpl) FindPrimes(req *proto.PrimeRequest, serverStream proto.AppService_FindPrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Finding prime number between %d and %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Sending prime number %d\n", no)
			serverStream.Send(&proto.PrimeResponse{
				PrimeNo: no,
			})
		}
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
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
