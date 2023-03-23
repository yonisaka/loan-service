package client

import (
	"context"

	pbUser "github.com/yonisaka/protobank/user"
)

// GetUser is a method
func (r GRPCClient) GetUser(ctx context.Context, payload *pbUser.UserByIDRequest) (*pbUser.UserResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	user, err := r.user.GetUser(ctx, payload)
	if err != nil {
		return nil, err
	}

	return user, nil
}