package main

import (
	"context"
	"fmt"
	proto "grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (asi *appServiceImpl) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var sum, count int32
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		no := req.GetNo()
		fmt.Printf("Receive %d for average calculation\n", no)
		sum += no
		count++
	}
	avg := sum / count
	res := &proto.AverageResponse{
		AvgResult: avg,
	}
	if err := serverStream.SendAndClose(res); err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (asi *appServiceImpl) Greet(serverStream proto.AppService_GreetServer) error {
	for {
		greetReq, err := serverStream.Recv()
		if code := status.Code(err); code == codes.Unavailable {
			fmt.Println("Client connection closed")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		person := greetReq.GetPerson()
		firstName := person.GetFirstName()
		lastName := person.GetLastName()
		log.Printf("Received greet request for %q and %q\n", firstName, lastName)
		message := fmt.Sprintf("Hi %s %s, Have a nice day!", firstName, lastName)
		time.Sleep(2 * time.Second)
		log.Printf("Sending response : %q\n", message)
		greetResp := &proto.GreetResponse{
			Message: message,
		}
		if err := serverStream.Send(greetResp); err != nil {
			if code := status.Code(err); code == codes.Unavailable {
				fmt.Println("Client connection closed")
				break
			}
		}
	}
	return nil
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
