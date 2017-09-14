package grpc

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/zerolog/log"
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/config"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/logger"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/search"
	"github.com/sh3rp/databox/secure"
	"github.com/sh3rp/databox/util"
)

type GRPCServer struct {
	Auth       auth.Authenticator
	TokenStore auth.TokenStore
	DB         db.BoxDB
	Search     search.SearchEngine
	Port       int
	HttpPort   int
	Filter     *secure.SecureFilter
}

func (s *GRPCServer) Start(certFile, keyFile string) {
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

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", s.HttpPort),
		Handler: http.HandlerFunc(handler),
	}
	go grpcServer.Serve(listener)
	go httpServer.ListenAndServeTLS(certFile, keyFile)
}

func (s *GRPCServer) Authenticate(ctx context.Context, req *msg.AuthRequest) (*msg.AuthResponse, error) {
	if s.Auth.Authenticate(req.Username, []byte(req.Password)) {
		log.Info().Msgf("User %s authenticated successfully, generating token", req.Username)
		token := s.TokenStore.GenerateToken(req.Username)
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

func (s *GRPCServer) UnlockBox(ctx context.Context, req *msg.UnlockRequest) (*msg.Box, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	box, err := s.DB.GetBoxById(*req.Box.Id)

	if bytes.Equal(secure.GetSignature([]byte(req.BoxPassword), box), box.EncryptedSignature) {
		s.Filter.UnlockBox(req.Token, box, []byte(req.BoxPassword))
	} else {
		return nil, errors.New("Incorrect box password")
	}

	return box, nil
}

func (s *GRPCServer) NewBox(ctx context.Context, req *msg.UnlockRequest) (*msg.Box, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	newBox, err := s.DB.NewBox(req.Box.Name, req.Box.Description, []byte(req.BoxPassword))

	if err != nil {
		logger.E(err)
		return nil, err
	}

	return newBox, nil
}

func (s *GRPCServer) SaveBox(ctx context.Context, req *msg.Request) (*msg.Box, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	err = s.DB.SaveBox(req.GetBox())

	if err != nil {
		logger.E(err)
		return nil, err
	}

	return req.GetBox(), err
}

func (s *GRPCServer) GetBoxById(ctx context.Context, req *msg.Request) (*msg.Box, error) {

	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	box, err := s.DB.GetBoxById(*req.GetBox().Id)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	return box, nil
}

func (s *GRPCServer) GetBoxes(ctx context.Context, req *msg.Request) (*msg.Boxes, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	boxes, err := s.DB.GetBoxes()

	if err != nil {
		logger.E(err)
		return nil, err
	}

	return &msg.Boxes{boxes}, nil
}

func (s *GRPCServer) NewLink(ctx context.Context, req *msg.Request) (*msg.Link, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	box, err := s.DB.GetBoxById(msg.Key{
		Type: msg.Key_BOX,
		Id:   req.GetLink().Id.BoxId,
	})

	if !s.Filter.IsUnlocked(req.Token, box) {
		return nil, errors.New(fmt.Sprintf("Box %s is locked", box.Id.Id))
	}

	if err != nil {
		logger.E(err)
		return nil, err
	}

	link, err := s.DB.NewLink(req.GetLink().Name, req.GetLink().Url, *box.Id)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	err = s.DB.SaveLink(s.Filter.EncryptLink(req.Token, link))

	if err != nil {
		logger.E(err)
		return nil, err
	}

	err = s.Search.Index(*link.Id, link.Tags)

	return link, err
}
func (s *GRPCServer) SaveLink(ctx context.Context, req *msg.Request) (*msg.Link, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	box, err := s.DB.GetBoxById(msg.Key{
		Type: msg.Key_BOX,
		Id:   req.GetLink().Id.BoxId,
	})

	if !s.Filter.IsUnlocked(req.Token, box) {
		return nil, errors.New(fmt.Sprintf("Box %s is locked", req.GetBox().Id.Id))
	}

	if err != nil {
		logger.E(err)
		return nil, err
	}

	err = s.DB.SaveLink(s.Filter.EncryptLink(req.Token, req.GetLink()))

	if err != nil {
		logger.E(err)
		return nil, err
	}

	err = s.Search.Index(*req.GetLink().Id, req.GetLink().Tags)

	return req.GetLink(), err
}
func (s *GRPCServer) GetLinkById(ctx context.Context, req *msg.Request) (*msg.Link, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	box, err := s.DB.GetBoxById(msg.Key{
		Type: msg.Key_BOX,
		Id:   req.GetLink().Id.BoxId,
	})

	if !s.Filter.IsUnlocked(req.Token, box) {
		return nil, errors.New(fmt.Sprintf("Box %s is locked", req.GetBox().Id.Id))
	}

	link, err := s.DB.GetLinkById(*req.GetLink().Id)

	return s.Filter.DecryptLink(req.Token, link), err
}
func (s *GRPCServer) GetLinksByBoxId(ctx context.Context, req *msg.Request) (*msg.Links, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	box, err := s.DB.GetBoxById(msg.Key{
		Type: msg.Key_BOX,
		Id:   req.GetBox().Id.Id,
	})

	if !s.Filter.IsUnlocked(req.Token, box) {
		return nil, errors.New(fmt.Sprintf("Box %s is locked", req.GetBox().Id.Id))
	}

	links, err := s.DB.GetLinksByBoxId(*req.GetBox().Id)

	var decryptedLinks []*msg.Link

	for _, l := range links {
		logger.OBJ(l)
		decryptedLinks = append(decryptedLinks, s.Filter.DecryptLink(req.Token, l))
	}

	return &msg.Links{decryptedLinks}, err
}
func (s *GRPCServer) SearchLinks(ctx context.Context, req *msg.Request) (*msg.Links, error) {
	err := s.TokenStore.ValidateToken(req.Token)

	if err != nil {
		logger.E(err)
		return nil, err
	}

	var links []*msg.Link
	linkIds := s.Search.Find(req.GetSearch().Term, int(req.GetSearch().Count), int(req.GetSearch().Page))

	for _, id := range linkIds {
		link, _ := s.DB.GetLinkById(id)
		box, err := s.DB.GetBoxById(msg.Key{
			Type: msg.Key_BOX,
			Id:   link.Id.BoxId,
		})

		if err != nil {
			logger.E(err)
		}

		if !s.Filter.IsUnlocked(req.Token, box) {
			logger.E(errors.New(fmt.Sprintf("Box %s is locked", req.GetBox().Id.Id)))
		}
		links = append(links, s.Filter.DecryptLink(req.Token, link))
	}
	return &msg.Links{links}, nil
}

func getTLSCredentials(certFile, keyFile string) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		logger.E(err)
		return nil, err
	}
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
	}), nil
}
