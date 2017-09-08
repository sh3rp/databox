package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/search"
	"github.com/sh3rp/databox/secure"
	"github.com/sh3rp/databox/server/grpc"
	"github.com/sh3rp/databox/server/web"
	"github.com/sh3rp/databox/server/web/routes"
)

type Server struct {
	HttpRouter *web.HttpServer
	GRPCRouter *grpc.GRPCServer
}

func NewServer(httpPort, grpcPort int, db db.BoxDB, search search.SearchEngine, a auth.Authenticator) *Server {
	tokenStore := auth.NewInMemoryTokenStore(20 * 60) // 20 minutes
	filter := secure.NewSecureFilter(a, tokenStore)

	return &Server{
		HttpRouter: &web.HttpServer{
			RouterBase: &routes.RouterBase{
				DB:         db,
				Search:     search,
				Auth:       a,
				TokenStore: tokenStore,
			},
			HttpRouter: gin.Default(),
			Port:       httpPort,
		},
		GRPCRouter: &grpc.GRPCServer{
			Auth:       a,
			TokenStore: tokenStore,
			DB:         db,
			Search:     search,
			Port:       grpcPort,
			Filter:     filter,
		},
	}
}

func (s *Server) Start() {
	go s.GRPCRouter.Start()
	go s.HttpRouter.Start()
}
