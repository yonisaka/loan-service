package client

import (
	"flag"
	"fmt"

	"github.com/yonisaka/loan-service/config"
	"github.com/yonisaka/loan-service/grpc/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGRPCUserService is a constructor
func NewGRPCUserService(config *config.Config) (*grpc.ClientConn, error) {
	flag.Parse()

	addr := flag.String("addrUserService", fmt.Sprintf("%s:%s", config.GRPCUserService.Host, config.GRPCUserService.Port), "The address to connect")
	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// NewGRPCBookService is a constructor
func NewGRPCBookService(config *config.Config) (*grpc.ClientConn, error) {
	flag.Parse()

	addr := flag.String("addrBookService", fmt.Sprintf("%s:%s", config.GRPCBookService.Host, config.GRPCBookService.Port), "The address to connect")
	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor.UnaryAuthClientInterceptor()),
		grpc.WithStreamInterceptor(interceptor.StreamAuthClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
