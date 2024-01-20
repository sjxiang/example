package main

import (
	"github.com/sjxiang/example/pb"
)

type Server struct {
	pb.UnimplementedBookstoreServer
	store Store
}

func NewServer(store Store) (*Server, error) {
	return &Server{
		store: store,
	}, nil 
}

