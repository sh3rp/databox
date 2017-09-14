package server

import (
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/search"
	"github.com/sh3rp/databox/secure"
	"github.com/sh3rp/databox/server/grpc"
	"github.com/sh3rp/databox/server/web"
)

type Server struct {
	HttpRouter *web.HttpServer
	GRPCRouter *grpc.GRPCServer
}

func NewServer(httpPort, grpcPort int, db db.BoxDB, search search.SearchEngine, a auth.Authenticator) *Server {
	tokenStore := auth.NewInMemoryTokenStore(20 * 60) // 20 minutes
	filter := secure.NewSecureFilter(a, tokenStore)

	return &Server{
		GRPCRouter: &grpc.GRPCServer{
			Auth:       a,
			TokenStore: tokenStore,
			DB:         db,
			Search:     search,
			Port:       grpcPort,
			Filter:     filter,
			HttpPort:   httpPort,
		},
	}
}

func (s *Server) Start(certFile, keyFile string) {
	s.GRPCRouter.Start(certFile, keyFile)
	//go s.HttpRouter.Start()
}
