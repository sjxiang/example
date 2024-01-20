package main

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	userAgentHeader = "user-agent"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (s *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	// 从 ctx 中获取 metadata
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}
		if p, ok := peer.FromContext(ctx); ok {
			mtdt.ClientIP = p.Addr.String()
		}	
	}

	return mtdt
}


/*

	md := s.extractMetadata(ctx)
	if md == nil {
		return nil, unauthenticatedError(errors.New("无效请求，丢失 metadata"))
	}

 */

