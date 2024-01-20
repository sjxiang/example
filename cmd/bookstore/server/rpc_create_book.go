package main

import (
	"context"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	
	"github.com/sjxiang/example/pb"
)


func (impl *Server) CreateBook(context.Context, *pb.CreateBookRequest) (*pb.Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
