package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/sjxiang/example/pb"
)


var EnableEncrypt = flag.Bool("enable", false, "开启 ssl/tls 加密传输")

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(authz("bearer sjxiang")))

	if *EnableEncrypt {		
		// 加载证书
		creds, err := credentials.NewClientTLSFromFile("./cert/server.crt", "")
		if err != nil {
			log.Fatalf("load failed, err: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts,  grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// 指定连接 server（写死）
	conn, err := grpc.Dial("127.0.0.1:8972", opts...)
	if err != nil {
		log.Fatalf("grpc.Dial failed, err:%v", err)
	}
	defer conn.Close()

	// 创建客户端
	c := pb.NewBookstoreClient(conn)
		
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.CreateShelf(ctx, &pb.CreateShelfRequest{
		Shelf: &pb.Shelf{
			Theme: "科普读物",
		},
	})
	
	if err != nil {
		// 封装了 FromError，更语义化一点
		st := status.Convert(err)
		switch st.Code() {
		case codes.InvalidArgument: 
			if len(st.Details()) > 0 {
				for _, detail := range st.Details() {
					switch t := detail.(type) {
					case *errdetails.BadRequest:
						log.Fatalf("请求参数错误导致失败：%v\n", t.GetFieldViolations())
					}
				}	
			}
		case codes.Internal:
			fmt.Println(st.Message())
		default:
			// 最差也能捞一个 codes.Unknown
		}

		log.Fatalf("其它失败，%v\n", err)
	}

	// 拿到了 gRPC 响应
	log.Fatalf("响应：%v\n", resp.String())
}


// 拦截器（客户端），认证
func authz(token string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 创建 metadata 1
		md := metadata.Pairs(
			"authorization", token,
		)

		// 创建 metadata 2
		// md ：= metadata.New(map[string]string{"authorization": "bear sjxiang-2023"})
		
		// 基于 metadata 创建 ctx
		ctx = metadata.NewOutgoingContext(ctx, md)

		// 拦截前
		return invoker(ctx, method, req, reply, cc, opts...)
	}	
}