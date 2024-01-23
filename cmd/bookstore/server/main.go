package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
