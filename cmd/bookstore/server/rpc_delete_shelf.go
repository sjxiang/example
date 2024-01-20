package main

import (
	"context"
	"errors"
	
	"gorm.io/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	
	"github.com/sjxiang/example/pb"
)


func (impl *Server) DeleteShelf(ctx context.Context, req *pb.DeleteShelfRequest) (*emptypb.Empty, error) {
	violations := validateDeleteShelfRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	_, err := impl.store.GetShelf(ctx, req.GetShelf())
	if err != nil {
		// 检查 ErrRecordNotFound 错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "invalid shelf id")  // 参数无效
		}

		return nil, status.Errorf(codes.Internal, "query failed: %s", err)
	}

	// 删之前，先查下有没有
	err = impl.store.DeleteShelf(ctx, req.GetShelf())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete failed: %s", err)
	}

	return &emptypb.Empty{}, nil
}


func validateDeleteShelfRequest(req *pb.DeleteShelfRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := ValidateShelfID(req.GetShelf()); err != nil {
		violations = append(violations, fieldViolation("shelf", err))
	}
	
	return violations
}
