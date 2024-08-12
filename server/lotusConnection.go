package server

import (
	"context"
	"errors"
	"log"
	"github/Alonza0314/lotus/packet"
	"reflect"
	"time"

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
	if _, err = stream.Read(buf); err != nil {
		log.Println("falied to read from stream:\n\t", err)
		return
	}

	req, err := packet.ParseRequest(buf)
	if err != nil {
		log.Println(err)
		return
	}

	value, err := callFunction(ls, req.Function, req.Parameters...)

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

func callFunction(ls lotusServer, function string, parameters ...interface{}) (interface{}, error) {
	f, ok := ls.serviceMap[function]
	if !ok {
		return nil, errors.New("failed to call function:\n\tno this function")
	}

	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		return nil, errors.New("failed to call function:\n\tno this function")
	}

	if len(parameters) != v.Type().NumIn() {
		return nil, errors.New("failed to call funtion:\n\tnumber of parameters is not correct")
	}

	reflectParameters := make([]reflect.Value, len(parameters))
	for i, param := range parameters {
		reflectParameters[i] = reflect.ValueOf(param)
	}

	result := v.Call(reflectParameters)

	var err error
	var values []interface{}

	for i, val := range result {
		if i == len(result) - 1 {
			if e, ok := val.Interface().(error); ok && e != nil {
				err = e
			}
		} else {
			values = append(values, val.Interface())
		}
	}
	return values, err
}
