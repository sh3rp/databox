package web

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/rs/zerolog/log"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/search"
)

type GRPCServer struct {
	DB     db.BoxDB
	Search search.SearchEngine
	Port   int
}

func (s *GRPCServer) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))

	if err != nil {
		log.Error().Msgf("GRPC server listen error: %v", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	msg.RegisterBoxServiceServer(grpcServer, s)
	grpcServer.Serve(listener)
}

func (s *GRPCServer) NewBox(ctx context.Context, box *msg.Box) (*msg.Box, error) {
	newBox, err := s.DB.NewBox(box.Name, box.Description)

	if err != nil {
		return nil, err
	}

	return newBox, nil
}

func (s *GRPCServer) SaveBox(ctx context.Context, box *msg.Box) (*msg.Box, error) {
	err := s.DB.SaveBox(box)

	if err != nil {
		return nil, err
	}

	return box, err
}

func (s *GRPCServer) GetBoxById(ctx context.Context, box *msg.Box) (*msg.Box, error) {

	box, err := s.DB.GetBoxById(box.Id)

	if err != nil {
		return nil, err
	}

	return box, nil
}

func (s *GRPCServer) GetBoxes(ctx context.Context, none *msg.None) (*msg.Boxes, error) {
	boxes, err := s.DB.GetBoxes()

	if err != nil {
		return nil, err
	}

	return &msg.Boxes{boxes}, nil
}

func (s *GRPCServer) NewLink(ctx context.Context, link *msg.Link) (*msg.Link, error) {
	_, err := s.DB.GetBoxById(link.BoxId)

	if err != nil {
		return nil, err
	}

	return s.DB.NewLink(link.Name, link.Url, link.BoxId)
}
func (s *GRPCServer) SaveLink(ctx context.Context, link *msg.Link) (*msg.Link, error) {
	_, err := s.DB.GetBoxById(link.BoxId)

	if err != nil {
		return nil, err
	}

	err = s.DB.SaveLink(link)

	return link, err
}
func (s *GRPCServer) GetLinkById(ctx context.Context, link *msg.Link) (*msg.Link, error) {
	return s.DB.GetLinkById(link.BoxId, link.Id)
}
func (s *GRPCServer) GetLinksByBoxId(ctx context.Context, box *msg.Box) (*msg.Links, error) {
	links, err := s.DB.GetLinksByBoxId(box.Id)

	return &msg.Links{links}, err
}
