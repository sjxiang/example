package main

import (
	"context"
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/sjxiang/example/pb"
)

const (
	defaultCursor   string = "0"  // 默认游标
	defaultPageSize int    = 2    // 默认每页显示数量
)
	

func (impl *Server) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	violations := validateListBooksRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// 默认，第一页
	var (
		cursor    = defaultCursor
		pageSize  = defaultPageSize
	)
	
	if len(req.GetPageToken()) > 0 {
		pageInfo := Token(req.GetPageToken()).Decode()
		
		cursor = pageInfo.NextID
		pageSize = int(pageInfo.PageSize)
	}

	// 查询 db，基于游标实现分页
	books, err := impl.store.GetBookListByShelfID(ctx, req.GetShelf(), cursor, pageSize+1)  // "超卖"
	if err == gorm.ErrEmptySlice {  // 没有数据（这里没用上，可见 gorm 一坨屎）
		return &pb.ListBooksResponse{}, nil 
	}
	if len(books) == 0 {
		return &pb.ListBooksResponse{}, nil
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "query failed: %s", err)
	}

	// 如果查询出来的结果比 pageSize 大，那么就说明有下一页
	var (
		hasNextPage   bool
		nextPageToken string
		realSize      int = len(books)
	)

	// 当查询数据库的结果数大于 pageSize，则有下一页
	if realSize > pageSize {
		hasNextPage = true
	}
	
	// 封装返回的数据（多的不要）
	rsp := convertBooksWithSize(books, pageSize)

	// 如果有下一页，就要生成下一页的 page_token
	if hasNextPage {
		nextID := fmt.Sprintf("%d", books[realSize-1].ID)  // 最后一个返回结果的 id  
		nextPageInfo := NewPage(nextID, int64(pageSize))
		nextPageToken = string(nextPageInfo.Encode())
	}

	return &pb.ListBooksResponse{
		Books:         rsp,
		NextPageToken: nextPageToken,
	}, nil
}

func validateListBooksRequest(req *pb.ListBooksRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := ValidateShelfID(req.GetShelf()); err != nil {
		violations = append(violations, fieldViolation("shelf", err))
	}
	if err :=  ValidatePageToken(req.GetPageToken()); err != nil {
		violations = append(violations, fieldViolation("page_token", err))
	}

	return violations
}




// hasPrev 若为 0，则无
// hasNext size + 1，判断