package test

import (
	"context"
	"lotus/server"
	"testing"
)

func TestServer(t *testing.T) {
	ls, err := server.NewLotusServer("test.pem")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	add := func(a, b int) int {
		return a + b
	}

	err = ls.RegisterService("add", add)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	ll, err := ls.Listen(":4433")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for i := 0; i < 1; i += 1 {
		lc, err := ll.Accept(context.Background())
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		go lc.HandleFunc(*ls)
	}
}
