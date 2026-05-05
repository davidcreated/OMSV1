package main

import (
	"log"
	"net/http"

	"github.com/OMSV1/common"
	pb "github.com/OMSV1/common/api"

	"github.com/joho/godotenv"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":3000")
	orderServiceAddr = common.EnvString("OMS_SERVICE_ADDR", "localhost:50051")
)

func main() {
	// what this does is to connect to the OMS service using gRPC and then start an HTTP server that will handle incoming requests and forward them to the OMS service
	conn, err := common.GrpcDial(orderServiceAddr)
	if err != nil {
		log.Fatal("Failed to connect to OMS service: ", err)
	}
	defer conn.Close()
	// the meaning of grpc is to create a client that can communicate with the OMS service using gRPC protocol, and the address of the OMS service is obtained from the environment variable OMS_SERVICE_ADDR, if it is not set, it will default to localhost:50051
	// grpc is a high performance, open source universal RPC framework that can run in any environment. It enables client and server applications to communicate transparently, and makes it easier to build connected systems.
	godotenv.Load()

	httpAddr := common.EnvString("HTTP_ADDR", ":3000")

	mux := http.NewServeMux()
	handler := NewHandler(pb.NewOmsServiceClient(conn))
	handler.registerRoutes(mux)

	log.Printf("Starting server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
