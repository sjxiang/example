package main

import (
	"context"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	
	"github.com/sjxiang/example/pb"
)

func (impl *Server) DeleteBook(context.Context, *pb.DeleteBookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
