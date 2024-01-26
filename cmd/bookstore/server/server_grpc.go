package main

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/sjxiang/example/pb"
)



func runGrpcServer(encrypt bool, store Store) {
	server, err := NewServer(store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(GrpcLogger))
	
	if encrypt {
		// 加载证书信息
		creds, err := credentials.NewServerTLSFromFile("./cert/server.crt", "./cert/server.key")  // 如果是当前文件夹 run，改为 cert/server.crt
		if err != nil {
			log.Fatal().Err(err).Msg("load failed")
		}
		opts = append(opts, grpc.Creds(creds))
	} 

	// 注册
	grpcSrv := grpc.NewServer(opts...)
	pb.RegisterBookstoreServer(grpcSrv, server)	

	// 监听
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8972))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())

	// 启动
	if err = grpcSrv.Serve(listener); err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}





// // 10.22.33.4:80/health   --> HTTP 200
// // 127.0.0.1:8972/health  --> gRPC ok
