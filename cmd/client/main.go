package main

import (
	"log"

	"github.com/markya0616/grpc-demo/api/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewBlockServiceClient(conn)
	response, err := c.GetBlock(context.Background(), &pb.GetBlockReq{Number: 1})
	if err != nil {
		log.Fatalf("Error when calling GetBlock: %s", err)
	}
	log.Printf("Response from server: %s", response.Block)

	// Streaming case
	stream, err := c.SubscribeBlock(context.Background(), &pb.SubscribeBlockReq{Number: 1})
	if err != nil {
		log.Fatalf("Error when calling SubscribeBlock: %s", err)
	}
	defer stream.CloseSend()

	for {
		event, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error when receve new blocks: %s", err)
		}
		log.Printf("Got block: %s", event.GetBlock())
	}
}
