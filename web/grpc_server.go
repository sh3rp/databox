package web

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/rs/zerolog/log"
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/config"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/search"
	"github.com/sh3rp/databox/util"
)

type GRPCServer struct {
	Auth       auth.Authenticator
	TokenStore auth.TokenStore
	DB         db.BoxDB
	Search     search.SearchEngine
	Port       int
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

func (s *GRPCServer) Authenticate(ctx context.Context, req *msg.AuthRequest) (*msg.AuthResponse, error) {
	if s.Auth.Authenticate(req.Username, req.Password) {
		log.Info().Msgf("User %s authenticated successfully, generating token", req.Username)
		token := s.TokenStore.GenerateToken(req.Username, time.Now().Add(20*time.Minute).UnixNano())
		log.Info().Msgf("Sending token: %v\n", token)
		return &msg.AuthResponse{
			Code:    0,
			Message: "ok",
			Token:   token,
		}, nil
	}
	log.Error().Msgf("User %s authentication FAILED", req.Username)
	return &msg.AuthResponse{
		Code:    1,
		Message: "User not authenticated",
	}, nil
}

func (s *GRPCServer) GetVersion(ctx context.Context, req *msg.Request) (*msg.Version, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		return nil, err
	}

	return util.GetVersion(), nil
}

func (s *GRPCServer) NewBox(ctx context.Context, req *msg.Request) (*msg.Box, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		log.Error().Msgf("NewBox: error validating token: %v", err)
		return nil, err
	}

	newBox, err := s.DB.NewBox(req.GetBox().Name, req.GetBox().Description)

	if err != nil {
		log.Error().Msgf("NewBox: error creating box: %v", err)
		return nil, err
	}

	return newBox, nil
}

func (s *GRPCServer) SaveBox(ctx context.Context, req *msg.Request) (*msg.Box, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		return nil, err
	}

	err = s.DB.SaveBox(req.GetBox())

	if err != nil {
		return nil, err
	}

	return req.GetBox(), err
}

func (s *GRPCServer) GetBoxById(ctx context.Context, req *msg.Request) (*msg.Box, error) {

	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		return nil, err
	}

	box, err := s.DB.GetBoxById(*req.GetBox().Id)

	if err != nil {
		return nil, err
	}

	return box, nil
}

func (s *GRPCServer) GetBoxes(ctx context.Context, req *msg.Request) (*msg.Boxes, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		return nil, err
	}

	boxes, err := s.DB.GetBoxes()

	if err != nil {
		return nil, err
	}

	return &msg.Boxes{boxes}, nil
}

func (s *GRPCServer) NewLink(ctx context.Context, req *msg.Request) (*msg.Link, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		return nil, err
	}

	box, err := s.DB.GetBoxById(msg.Key{
		Type: msg.Key_BOX,
		Id:   req.GetLink().Id.BoxId,
	})

	if err != nil {
		return nil, err
	}

	link, err := s.DB.NewLink(req.GetLink().Name, req.GetLink().Url, *box.Id)

	if err != nil {
		return nil, err
	}

	err = s.Search.Index(*link.Id, link.Tags)

	return link, err
}
func (s *GRPCServer) SaveLink(ctx context.Context, req *msg.Request) (*msg.Link, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		return nil, err
	}

	_, err = s.DB.GetBoxById(msg.Key{
		Type: msg.Key_BOX,
		Id:   req.GetLink().Id.BoxId,
	})

	if err != nil {
		return nil, err
	}

	err = s.DB.SaveLink(req.GetLink())

	if err != nil {
		return nil, err
	}

	err = s.Search.Index(*req.GetLink().Id, req.GetLink().Tags)

	return req.GetLink(), err
}
func (s *GRPCServer) GetLinkById(ctx context.Context, req *msg.Request) (*msg.Link, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		return nil, err
	}

	return s.DB.GetLinkById(*req.GetLink().Id)
}
func (s *GRPCServer) GetLinksByBoxId(ctx context.Context, req *msg.Request) (*msg.Links, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		return nil, err
	}

	links, err := s.DB.GetLinksByBoxId(*req.GetBox().Id)

	return &msg.Links{links}, err
}
func (s *GRPCServer) SearchLinks(ctx context.Context, req *msg.Request) (*msg.Links, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		return nil, err
	}

	var links []*msg.Link
	linkIds := s.Search.Find(req.GetSearch().Term, int(req.GetSearch().Count), int(req.GetSearch().Page))

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
