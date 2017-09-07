package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/search"
)

type Server struct {
	HttpRouter *HttpServer
	GRPCRouter *GRPCServer
	DB         db.BoxDB
	Search     search.SearchEngine
	HttpPort   int
	GRPCPort   int
}

func NewServer(httpPort, grpcPort int, db db.BoxDB, search search.SearchEngine, a auth.Authenticator) *Server {
	return &Server{
		HttpRouter: &HttpServer{DB: db, Search: search, HttpRouter: gin.Default(), Port: httpPort},
		GRPCRouter: &GRPCServer{Auth: a, TokenStore: auth.NewInMemoryTokenStore(), DB: db, Search: search, Port: grpcPort},
		DB:         db,
		Search:     search,
		HttpPort:   httpPort,
		GRPCPort:   grpcPort,
	}
}

func (s *Server) Start() {
	go s.GRPCRouter.Start()
	go s.HttpRouter.Start()
}
