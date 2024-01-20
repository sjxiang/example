package main

import (
	"context"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"github.com/go-sql-driver/mysql"
	
	"github.com/sjxiang/example/pb"
)

func (impl *Server) CreateShelf(ctx context.Context, req *pb.CreateShelfRequest) (*pb.Shelf, error) {
	// 校验请求字段
	violations := validateCreateShelfRequest(req)
	// 至少有一个无效字段
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	data := Shelf{
		Theme: req.GetShelf().GetTheme(),
		Size:  req.GetShelf().GetSize(),
	}

	shelf, err := impl.store.CreateShelf(ctx, data)
	if err != nil {
		// 违反唯一索引约束 Error 1062 (23000): Duplicate entry
		const uniqueViolation uint16 = 1062  

		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case uniqueViolation: 
				return nil, status.Errorf(codes.AlreadyExists, err.Error())
			}
		}
		return nil, status.Errorf(codes.Internal, "create failed: %s", err)
	}

	rsp := convertShelf(shelf)

	return rsp, nil
}

func validateCreateShelfRequest(req *pb.CreateShelfRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := ValidateShelfTheme(req.GetShelf().GetTheme()); err != nil {
		violations = append(violations, fieldViolation("theme", err))
	}
	
	return violations
}
