package main

import (
	"context"
	"fmt"
	proto "grpc-app/proto"
	"log"

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

	doRequestResponse(ctx, appServiceClient)

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
