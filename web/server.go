package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/msg"
)

type Server struct {
	HttpRouter *gin.Engine
	GRPCRouter *GRPCServer
	DB         db.BoxDB
	HttpPort   int
	GRPCPort   int
}

func NewServer(httpPort, grpcPort int, db db.BoxDB) *Server {
	return &Server{
		HttpRouter: gin.Default(),
		GRPCRouter: &GRPCServer{db},
		DB:         db,
		HttpPort:   httpPort,
		GRPCPort:   grpcPort,
	}
}

func (s *Server) Start() {
	db := s.DB
	r := s.HttpRouter
	r.GET("/box", func(c *gin.Context) {
		boxes, err := db.GetBoxes()
		if err != nil {
			c.JSON(200, Error(E_DB_BOX_NOT_FOUND, err))
		} else {
			c.JSON(200, Success(boxes))
		}
	})
	r.POST("/box", func(c *gin.Context) {
		var box msg.Box
		err := c.BindJSON(&box)
		if err == nil {
			if box.Id == "" {
				newBox, err := db.NewBox(box.Name, box.Description, false)
				if err != nil {
					c.JSON(200, Error(E_DB_CREATE_BOX, err))
					return
				} else {
					c.JSON(200, Success(newBox))
					return
				}
			}
			c.JSON(200, Success(box))
		} else {
			c.JSON(200, Error(E_IO_DECODE_BOX, err))
		}
	})
	r.GET("/box/:id", func(c *gin.Context) {
		id := c.Param("id")

		box, err := s.DB.GetBoxById(id)

		if err == nil {
			c.JSON(200, Success(box))
		} else {
			c.JSON(200, Error(E_DB_BOX_NOT_FOUND, err))
		}

	})
	r.GET("/box/:id/link", func(c *gin.Context) {
		id := c.Param("id")

		links, err := db.GetLinksByBoxId(id)

		if err == nil {
			c.JSON(200, Success(links))
		} else {
			c.JSON(200, Error(E_DB_LINK_NOT_FOUND, err))
		}

	})
	r.POST("/box/:id/link", func(c *gin.Context) {
		boxId := c.Param("boxId")

		var linkUpdate msg.Link

		err := c.BindJSON(&linkUpdate)

		if err != nil {
			c.JSON(200, Error(E_IO_DECODE_LINK, err))
			return
		}

		_, err = s.DB.GetBoxById(boxId)

		if err != nil {
			c.JSON(200, Error(E_DB_BOX_NOT_FOUND, err))
			return
		}

		link, err := s.DB.GetLinkById(linkUpdate.Id)

		if err != nil {
			link, err = s.DB.NewLink(linkUpdate.Name, linkUpdate.Url, boxId)

			if err != nil {
				c.JSON(200, Error(E_DB_CREATE_LINK, err))
				return
			} else {
				c.JSON(200, Success(link))
			}
		} else {
			linkUpdate.BoxId = boxId
			err = s.DB.SaveLink(&linkUpdate)
			if err != nil {
				c.JSON(200, Error(E_DB_UPDATE_LINK, err))
				return
			} else {
				c.JSON(200, Success(linkUpdate))
			}
		}

	})
	r.GET("/box/:id/link/:linkId", func(c *gin.Context) {
		boxId := c.Param("id")

		_, err := s.DB.GetBoxById(boxId)

		if err != nil {
			c.JSON(200, Error(E_DB_BOX_NOT_FOUND, err))
			return
		}

		linkId := c.Param("linkId")

		link, err := s.DB.GetLinkById(linkId)

		if err != nil {
			c.JSON(200, Error(E_DB_LINK_NOT_FOUND, err))
			return
		} else {
			c.JSON(200, Success(link))
		}
	})
	r.GET("/box/:id/image", func(c *gin.Context) {

	})
	r.POST("/box/:id/image/:imageId", func(c *gin.Context) {

	})
	go s.GRPCRouter.Start(s.GRPCPort)
	r.Run(fmt.Sprintf(":%d", s.HttpPort))
}
