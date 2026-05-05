package common

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GrpcDial(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
