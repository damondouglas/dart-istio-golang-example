package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	echo "temp/proto"
)

const (
	portKey = "PORT"
)

var (
	address string
)

func init() {
	port := os.Getenv(portKey)
	if port == "" {
		log.Fatalf("%s empty but expected from environment variables", portKey)
	}
	address = fmt.Sprintf(":%s", port)
}

type server struct {}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &server{})
	log.Printf("Listening on %s ...", address)
	log.Fatal(s.Serve(lis))
}

func (*server) Get(ctx context.Context, in *echo.GetRequest) (response *echo.GetResponse, err error) {
	log.Printf("Request received %v", in)
	response = &echo.GetResponse{
		Payload: in.Payload + in.Payload,
	}
	return
}