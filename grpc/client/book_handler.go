package client

import (
	"context"

	pbBook "github.com/yonisaka/protobank/book"
)

// GetBook is a method
func (r GRPCClient) GetBook(ctx context.Context, payload *pbBook.BookByIDRequest) (*pbBook.BookResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	book, err := r.book.GetBook(ctx, payload)
	if err != nil {
		return nil, err
	}

	return book, nil
}
