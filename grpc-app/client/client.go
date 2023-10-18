package main

import (
	"context"
	"fmt"
	proto "grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// create a grpc client connection
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalln(err)
	}

	//create an instance of the service proxy
	appServiceClient := proto.NewAppServiceClient(clientConn)

	ctx := context.Background()

	// doRequestResponse(ctx, appServiceClient)
	// doServerStreaming(ctx, appServiceClient)
	// doClientStreaming(ctx, appServiceClient)
	doBiDiStreaming(ctx, appServiceClient)

}

func doRequestResponse(ctx context.Context, appServiceClient proto.AppServiceClient) {
	// request & response
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}

	addResponse, err := appServiceClient.Add(ctx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(addResponse.GetResult())
}

func doServerStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	primeReq := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	clientStream, err := appServiceClient.FindPrimes(ctx, primeReq)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		primeRes, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All the prime numbers are received")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		primeNo := primeRes.GetPrimeNo()
		fmt.Printf("prime no : %d\n", primeNo)
	}
}

func doClientStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	data := []int32{3, 1, 4, 2, 8, 6, 7, 9, 5}
	clientStream, err := appServiceClient.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, val := range data {
		req := &proto.AverageRequest{
			No: val,
		}
		fmt.Printf("Average req, no : %d\n", val)
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Average : %d\n", res.GetAvgResult())
}

func doBiDiStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	clientStream, err := appServiceClient.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	go sendRequests(ctx, clientStream)
	done := make(chan struct{})
	go func() {
		fmt.Println("Press ENTER to cancel")
		fmt.Scanln()
		clientStream.CloseSend()
		close(done)
	}()
	go recvResponse(ctx, clientStream)
	// return done
	<-done
}

func sendRequests(ctx context.Context, clientStream proto.AppService_GreetClient) {
	persons := []*proto.PersonName{
		{FirstName: "Magesh", LastName: "Kuppan"},
		{FirstName: "Suresh", LastName: "Kannan"},
		{FirstName: "Ramesh", LastName: "Jayaraman"},
		{FirstName: "Rajesh", LastName: "Pandit"},
		{FirstName: "Ganesh", LastName: "Kumar"},
	}

	// done := make(chan struct{})

	for _, person := range persons {
		req := &proto.GreetRequest{
			Person: person,
		}
		log.Printf("Sending Person : %s %s\n", person.FirstName, person.LastName)
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func recvResponse(ctx context.Context, clientStream proto.AppService_GreetClient) {
	for {
		res, err := clientStream.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(res.GetMessage())
	}
}
