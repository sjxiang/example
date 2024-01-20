package main

import (
	"context"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	
	"github.com/sjxiang/example/pb"
)


func (impl *Server) GetBook(context.Context, *pb.GetBookRequest) (*pb.Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
