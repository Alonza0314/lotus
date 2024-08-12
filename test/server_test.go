package test

import (
	"context"
	"testing"
	"time"

	"github.com/Alonza0314/lotus/server"
)

func TestServer(t *testing.T) {
	lserver, err := server.NewLotusServer("test.pem")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	add := func(a, b interface{}) float64 {
		return a.(float64) + b.(float64)
	}

	err = lserver.RegisterService("add", add)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	llistener, err := lserver.Listen(":4433")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for i := 0; i < 1; i += 1{
		lconn, err := llistener.Accept(context.Background())
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		go lconn.HandleFunc(*lserver)
	}

	time.Sleep(1 * time.Second)
}
