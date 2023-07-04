package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	proto "week3_docker/pkg/api/contact_service"
)

var id *uint64

func init() {
	id = flag.Uint64("id", 0, "unsub account id")
}

func main() {
	flag.Parse()

	if id == nil {
		fmt.Println("require account id")
	}

	req := &proto.UnsubAccountRequest{
		Id: *id,
	}

	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //nolint
	if err != nil {
		log.Fatal("connection grpc server error")
	}
	client := proto.NewContactServiceClient(con)
	_, err = client.UnsubAccount(context.Background(), req)
	if err != nil {
		log.Printf("error grpc.UnsubAccount: %v", err)
	} else {
		log.Println("success delete acccount ", *id)
	}
}
