package server

import (
	"context"
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/Alonza0314/lotus/packet"

	"github.com/quic-go/quic-go"
)

type lotusConnection struct {
	connection quic.Connection
}

func NewLotusConnection(connection quic.Connection) (*lotusConnection, error) {
	return &lotusConnection{connection: connection}, nil
}

func (lc *lotusConnection) HandleFunc(ls lotusServer) {
	stream, err := lc.connection.AcceptStream(context.Background())
	if err != nil {
		log.Println("failed to accept stream:\n\t", err)
		return
	}
	defer stream.Close()

	buf := make([]byte, 4096)
	n, err := stream.Read(buf)
	if err != nil {
		log.Println("falied to read from stream:\n\t", err)
		return
	}

	req, err := packet.ParseRequest(buf[:n])
	if err != nil {
		log.Println(err)
		return
	}

	value, err := callFunction(ls, req.Function, req.Args...)

	var res *packet.Response
	if err != nil {
		res, _ = packet.NewResponse("fail", req.Function, value.([]interface{}), err.Error())
	} else {
		res, _ = packet.NewResponse("success", req.Function, value.([]interface{}), "")
	}

	reply, err := res.MakeJson()
	if err != nil {
		log.Println(err)
	}

	_, err = stream.Write(reply)
	if err != nil {
		log.Println("failed to write to stream:", err)
		return
	}

	time.Sleep(1 * time.Second)
}

func callFunction(ls lotusServer, function string, args ...interface{}) (interface{}, error) {
	f, ok := ls.serviceMap[function]
	if !ok {
		return nil, errors.New("failed to call function:\n\tno this function")
	}

	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		return nil, errors.New("failed to call function:\n\tno this function")
	}

	if len(args) != v.Type().NumIn() {
		return nil, errors.New("failed to call funtion:\n\tnumber of args is not correct")
	}

	reflectArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		reflectArgs[i] = reflect.ValueOf(arg)
	}

	result := v.Call(reflectArgs)

	var values []interface{}

	for _, val := range result {
		values = append(values, val.Interface())
	}
	return values, nil
}
