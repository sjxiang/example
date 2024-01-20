package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/sjxiang/example/pb"
)

const (
	dataSourceUrl = "root:my-secret-pw@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=true&loc=Local"
)

var (
	encrypt = flag.Bool("encrypt", false, "开启 ssl/tls 加密传输")
)

func main() {
	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})  // 日志输出编码：console 或 json

	db, err := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := NewStore(db)

	// http
	go runGatewayServer(store)

	// grpc
	runGrpcServer(*encrypt, store)
}


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


func runGatewayServer(store Store) {

	server, err := NewServer(store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	// 请求参数，驼峰转下划线
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterBookstoreHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	
	// 中间件
	handler := HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}
}


/*
	艹，为啥这么多写法，烦人

	srv := &http.Server{
		Addr:    ":8090",
		Handler: HttpLogger(grpcMux),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}

*/
