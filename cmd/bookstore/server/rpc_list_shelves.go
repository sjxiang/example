package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	"github.com/sjxiang/example/pb"
)


func (impl *Server) ListShelves(ctx context.Context, req *emptypb.Empty) (*pb.ListShelvesResponse, error) {

	shelves, err := impl.store.ListShelves(ctx)
	if err == gorm.ErrEmptySlice {  // 没有数据
		return &pb.ListShelvesResponse{}, nil 
	}
	if err != nil {  // 查询数据库失败
		return nil, status.Errorf(codes.Internal, "query failed: %s", err)
	}
	
	rsp := &pb.ListShelvesResponse{
		Shelves: convertShelves(shelves),
	}

	return rsp, nil 
}