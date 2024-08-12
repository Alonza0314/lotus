package client

import (
	"context"
	"errors"
	"log"

	"github.com/Alonza0314/lotus/packet"
	"github.com/quic-go/quic-go"
)

type lotusConnection struct {
	connection quic.Connection
}

func NewLotusConnection(connection quic.Connection) (*lotusConnection, error) {
	return &lotusConnection{connection: connection}, nil
}

func (lc *lotusConnection) Close() {
	if err := lc.connection.CloseWithError(0, ""); err != nil {
		log.Fatalln("failed to cloes connection:\n\t", err.Error())
	}
}

func (lc *lotusConnection) Call(ctx context.Context, function string, args []interface{}, reply *[]interface{}) error {
	req, err := packet.NewRequest(function, args)
	if err != nil {
		return err
	}
	reqJsonByte, err := req.MakeJson()
	if err != nil {
		return err
	}

	stream, err := lc.connection.OpenStreamSync(ctx)
	if err != nil {
		return errors.New("fail to open stream:\n\t" + err.Error())
	}
	defer stream.Close()

	_, err = stream.Write(reqJsonByte)
	if err != nil {
		return errors.New("failed to write to stream:\n\t" + err.Error())
	}

	buf := make([]byte, 4096)
	n, err := stream.Read(buf)
	if err != nil {
		return errors.New("failed to read from stream:\n\t" + err.Error())
	}

	res, err := packet.ParseResponse(buf[:n])
	if err != nil {
		return err
	}

	if res.Condition == "fail" {
		return errors.New(res.ErrorMsg)
	}

	if res.Function != function {
		return errors.New("failed to get correct response from function:\n\treturn value is not correspond to the function")
	}

	*reply = make([]interface{}, len(res.Replys))
	copy(*reply, res.Replys)

	return nil
}
