package server

import (
	"crypto/tls"
	"errors"
	"lotus/security"
	"reflect"

	"github.com/quic-go/quic-go"
)

type lotusServer struct {
	tlsConfig  *tls.Config
	quicConfig *quic.Config
	serviceMap map[string]interface{}
}

func NewLotusServer(pemPath string) (*lotusServer, error) {
	cert, err := security.LoadTLSCertificate(pemPath)
	if err != nil {
		return nil, errors.New("failed to new lotus server:\n\t" + err.Error())
	}
	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS13,
		Certificates: []tls.Certificate{*cert},
	}
	return &lotusServer{
		tlsConfig:  tlsConfig,
		quicConfig: nil,
		serviceMap: make(map[string]interface{}),
	}, nil
}

func (ls *lotusServer) Listen(addr string) (*lotusListener, error) {
	listener, err := quic.ListenAddr(addr, ls.tlsConfig, ls.quicConfig)
	if err != nil {
		return nil, errors.New("failed to listen:\n\t" + err.Error())
	}
	return NewLotusListener(listener)
}

func (ls *lotusServer) RegisterService(name string, function interface{}) error {
	if name == "" {
		return errors.New("failed to register service:\n\tservice name is empty")
	}
	if function == nil {
		return errors.New("failed to register service:\n\tfunction is nil")
	}

	funcType := reflect.TypeOf((function))
	if funcType.Kind() != reflect.Func {
		return errors.New("failed to register service:\n\tfunction parameters is not function type")
	}

	ls.serviceMap[name] = function

	return nil
}