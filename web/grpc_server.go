package web

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/rs/zerolog/log"
	"github.com/sh3rp/databox/config"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/search"
	"github.com/sh3rp/databox/util"
)

type GRPCServer struct {
	DB     db.BoxDB
	Search search.SearchEngine
	Port   int
}

func (s *GRPCServer) Start() {
	serverConfig := &config.ServerConfig{}
	serverConfig.Read("server.json")

	credentials, err := getTLSCredentials(serverConfig.CertFile, serverConfig.KeyFile)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))

	if err != nil {
		log.Error().Msgf("GRPC server listen error: %v", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(grpc.Creds(credentials))

	msg.RegisterBoxServiceServer(grpcServer, s)
	grpcServer.Serve(listener)
}

func (s *GRPCServer) GetVersion(ctx context.Context, none *msg.None) (*msg.Version, error) {
	return util.GetVersion(), nil
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

	box, err := s.DB.GetBoxById(*box.Id)

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
	box, err := s.DB.GetBoxById(msg.Key{
		Type: msg.Key_BOX,
		Id:   link.Id.BoxId,
	})

	if err != nil {
		return nil, err
	}

	link, err = s.DB.NewLink(link.Name, link.Url, *box.Id)

	if err != nil {
		return nil, err
	}

	err = s.Search.Index(*link.Id, link.Tags)

	return link, err
}
func (s *GRPCServer) SaveLink(ctx context.Context, link *msg.Link) (*msg.Link, error) {
	_, err := s.DB.GetBoxById(msg.Key{
		Type: msg.Key_BOX,
		Id:   link.Id.BoxId,
	})

	if err != nil {
		return nil, err
	}

	err = s.DB.SaveLink(link)

	if err != nil {
		return nil, err
	}

	err = s.Search.Index(*link.Id, link.Tags)

	return link, err
}
func (s *GRPCServer) GetLinkById(ctx context.Context, link *msg.Link) (*msg.Link, error) {
	return s.DB.GetLinkById(*link.Id)
}
func (s *GRPCServer) GetLinksByBoxId(ctx context.Context, box *msg.Box) (*msg.Links, error) {
	links, err := s.DB.GetLinksByBoxId(*box.Id)

	return &msg.Links{links}, err
}
func (s *GRPCServer) SearchLinks(ctx context.Context, search *msg.Search) (*msg.Links, error) {
	var links []*msg.Link
	linkIds := s.Search.Find(search.Term, int(search.Count), int(search.Page))

	for _, id := range linkIds {
		link, _ := s.DB.GetLinkById(id)
		links = append(links, link)
	}
	return &msg.Links{links}, nil
}

func getTLSCredentials(certFile, keyFile string) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
	}), nil
}
