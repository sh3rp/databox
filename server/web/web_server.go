package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/server/web/routes"
)

type HttpServer struct {
	RouterBase *routes.RouterBase
	HttpRouter *gin.Engine
	Port       int
}

func (s *HttpServer) Start() {
	r := s.HttpRouter

	r.GET("/box", s.RouterBase.GetBox)
	r.GET("/box/:id", s.RouterBase.GetBoxById)
	r.GET("/box/:id/link", s.RouterBase.GetLink)
	r.GET("/box/:id/link/:linkId", s.RouterBase.GetLinkById)

	r.POST("/box", s.RouterBase.PostBox)
	r.POST("/box/:id/link", s.RouterBase.PostLink)

	r.Run(fmt.Sprintf(":%d", s.Port))
}
