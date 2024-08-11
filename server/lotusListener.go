package server

import (
	"context"
	"errors"

	"github.com/quic-go/quic-go"
)

type lotusListener struct {
	listener *quic.Listener
}

func NewLotusListener(listener *quic.Listener) (*lotusListener, error) {
	return &lotusListener{listener: listener}, nil
}

func (ll *lotusListener) Accept(ctx context.Context) (*lotusConnection, error) {
	conn, err := ll.listener.Accept(ctx)
	if err != nil {
		return nil, errors.New("failed to accept:\n\t" + err.Error())
	}
	return NewLotusConnection(conn)
}
