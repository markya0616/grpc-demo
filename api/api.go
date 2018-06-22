package api

import (
	"fmt"
	"log"
	"time"

	"github.com/markya0616/grpc-demo/api/pb"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// GetBlock returns the block
func (s *Server) GetBlock(ctx context.Context, req *pb.GetBlockReq) (*pb.GetBlockRsp, error) {
	log.Printf("Receive block %v", req.Number)
	return &pb.GetBlockRsp{
		Block: fmt.Sprintf("Block %v", req.Number),
	}, nil
}

// SubscribeBlock subscribes new block events
func (s *Server) SubscribeBlock(req *pb.SubscribeBlockReq, server pb.BlockService_SubscribeBlockServer) error {
	ticker := time.NewTicker(1000 * time.Millisecond).C
	count := req.Number
	for {
		select {
		case <-ticker:
			count++
			err := server.Send(&pb.SubscribeBlockRsp{
				Block: fmt.Sprintf("Stream block %v", count),
			})
			if err != nil {
				return err
			}
			log.Printf("Send block %v!", count)
		case <-server.Context().Done():
			log.Printf("Context done!")
			return nil
		}
	}
}
