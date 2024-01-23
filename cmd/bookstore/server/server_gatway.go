package main

import (
	"context"
	"net"
	"net/http"
	
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
	
	"github.com/sjxiang/example/pb"
)

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
