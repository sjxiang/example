package main

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/hashicorp/consul/api"

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

	// 给我们的 gRPC 服务增加了健康检查的处理逻辑
	healthCheck := health.NewServer()
	healthpb.RegisterHealthServer(grpcSrv, healthCheck)  // consul 发来健康检查的RPC请求，这个负责返回OK

	// 连接至 consul
	cc, err := api.NewClient(api.DefaultConfig()) 
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to consul")
	}

	// 获取本机的出口 ip
	ip, err := GetOutboundIP()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot get local ip")
	}

	log.Info().Msgf("local ip: %s", ip)
	
	// 将我们的 gRPC 服务注册到 consul
	// 1. 定义我们的服务
	// 配置健康检查策略，告诉 consul 如何进行健康检查
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", ip, 8972), // 外网地址，相对于容器内的地址
		Timeout:                        "5s",
		Interval:                       "5s",  // 健康检查，间隔
		DeregisterCriticalServiceAfter: "10s", // 10 秒钟后，注销掉不健康的服务节点
	}

	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", "bookstore", ip, 8972), // 服务唯一 ID
		Name:    "bookstore",
		Tags:    []string{"xxx"},
		Address: ip,
		Port:    8972,
		Check:   check,
	}
	
	// 2. 注册服务到 consul
	if err := cc.Agent().ServiceRegister(srv); err != nil {
		log.Fatal().Err(err).Msg("cannot register with consul")
	}


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
