package test

import (
	"context"
	"testing"

	"github.com/Alonza0314/lotus/client"
)

func TestClient(t *testing.T) {
	lclient, err := client.NewLotusClient(":4433", true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	lconn, err := lclient.Dial(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	defer lconn.Close()

	function, args, reply := "add", []interface{}{1, 2}, []interface{}{}

	if err := lconn.Call(context.Background(), function, args, &reply); err != nil {
		t.Fatal(err)
	}

	if reply[0].(float64) != 3 {
		t.Fatal("expected to get 3, got", reply[0])
	}
}
