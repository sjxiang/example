package main

import (
	"context"
	"errors"
	
	"gorm.io/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	
	"github.com/sjxiang/example/pb"
)


func (impl *Server) GetShelf(ctx context.Context, req *pb.GetShelfRequest) (*pb.Shelf, error) {

	violations := validateGetShelfRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	shelf, err := impl.store.GetShelf(ctx, req.GetShelf())
	if err != nil {
		// 检查 ErrRecordNotFound 错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "invalid shelf id")
		}

		return nil, status.Errorf(codes.Internal, "query failed: %s", err)
	}
	
	rsp := convertShelf(shelf)

	return rsp, nil 
}


func validateGetShelfRequest(req *pb.GetShelfRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := ValidateShelfID(req.GetShelf()); err != nil {
		violations = append(violations, fieldViolation("shelf", err))
	}
	
	return violations
}

