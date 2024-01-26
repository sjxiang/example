package main

import (
	"context"

	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/metadata"
)


const (
	objectWildcard = "*"           // 所有资源
	produceAction  = "produce"     // 操作
	consumeAction  = "consume"
)


const (
	userAgentHeader = "user-agent"
	subjectHeader   = "subject"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
	Subject   string
}

func (s *Server) extractMetadataa(ctx context.Context) *Metadata {

	mtdt := &Metadata{}

	// 元信息
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}
		if sujects := md.Get(subjectHeader); len(sujects) > 0 {
			mtdt.Subject = sujects[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}
	
	return mtdt
}


/*

	md := s.extractMetadata(ctx)
	if md == nil {
		return nil, unauthenticatedError(errors.New("无效请求，丢失 metadata"))
	}

	if err := s.Authorizer.Authorize(md.Subject, objectWildcard, produceAction); err != nil {
		return nil, err
	}

 */