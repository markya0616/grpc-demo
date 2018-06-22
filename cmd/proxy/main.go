package main

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/markya0616/grpc-demo/api/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	mux := runtime.NewServeMux()
	err := pb.RegisterBlockServiceHandlerFromEndpoint(context.Background(), mux, ":7777", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("Failed to register rest server: %s", err)
	}

	for {
		log.Printf("starting HTTP/1.1 REST server on :5454")
		http.ListenAndServe(":5454", mux)
	}
}
