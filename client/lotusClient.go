package client

import (
	"context"
	"crypto/tls"
	"errors"

	"github.com/quic-go/quic-go"
)

type lotusClient struct {
	addr       string
	tlsConfig  *tls.Config
	quicConfig *quic.Config
}

func NewLotusClient(addr string, insecureSkipVerify bool) (*lotusClient, error) {
	if addr == "" {
		return nil, errors.New("failed to new lotus client:\n\taddr can't be empty")
	}
	return &lotusClient{
		addr: addr,
		tlsConfig: &tls.Config {
			InsecureSkipVerify: insecureSkipVerify,
		},
		quicConfig: nil,
	}, nil
}

func (lc *lotusClient) Dial(ctx context.Context) (*lotusConnection, error) {
	conn, err := quic.DialAddr(ctx, lc.addr, lc.tlsConfig, lc.quicConfig)
	if err != nil {
		return nil, errors.New("failed to dial:\n\t" + err.Error())
	}
	return NewLotusConnection(conn)
}