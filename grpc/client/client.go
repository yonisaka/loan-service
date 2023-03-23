package client

import (
	pbAuth "github.com/yonisaka/protobank/auth"
	pbBook "github.com/yonisaka/protobank/book"
	pbLog "github.com/yonisaka/protobank/log"
	pbUser "github.com/yonisaka/protobank/user"
	"google.golang.org/grpc"
)

// GRPCClient is a struct
type GRPCClient struct {
	auth    pbAuth.AuthClient
	httpLog pbLog.LogServiceClient
	book    pbBook.BookServiceClient
	user pbUser.UserServiceClient
}

// NewGRPCClient is constructor
func NewGRPCClient(userServiceConn grpc.ClientConnInterface, bookServiceConn grpc.ClientConnInterface) *GRPCClient {
	return &GRPCClient{
		auth:    pbAuth.NewAuthClient(userServiceConn),
		httpLog: pbLog.NewLogServiceClient(userServiceConn),
		book:    pbBook.NewBookServiceClient(bookServiceConn),
		user: pbUser.NewUserServiceClient(userServiceConn),
	}
}
