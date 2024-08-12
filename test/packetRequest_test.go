package test

import (
	"github.com/Alonza0314/lotus/packet"
	"testing"
)

func TestNewRequest(t *testing.T) {
	function := "testFunction"
	args := []interface{}{1, "arg"}
	req, err := packet.NewRequest(function, args)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if req.Function != function {
		t.Errorf("expected function %v, got %v", function, req.Function)
	}
	if len(req.Args) != len(args) {
		t.Errorf("expected args length %v, got %v", len(args), len(req.Args))
	}

	_, err = packet.NewRequest("", args)
	if err == nil {
		t.Fatal("expected error for empty function name, got none")
	}
}




func TestParseRequest(t *testing.T) {
	msg := `{"function":"testFunction","args":[1,"arg"]}`
	expectedFunction := "testFunction"
	expectedArgs := []interface{}{1.0, "arg"}

	req, err := packet.ParseRequest([]byte(msg))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if req.Function != expectedFunction {
		t.Errorf("expected function %v, got %v", expectedFunction, req.Function)
	}
	if len(req.Args) != len(expectedArgs) {
		t.Errorf("expected args length %v, got %v", len(expectedArgs), len(req.Args))
	}
	for i, param := range req.Args {
		if param != expectedArgs[i] {
			t.Errorf("expected arg %v at index %d, got %v", expectedArgs[i], i, param)
		}
	}

	invalidMsg := `{"function":"testFunction","args":[1,"arg"]`
	_, err = packet.ParseRequest([]byte(invalidMsg))
	if err == nil {
		t.Fatal("expected error for invalid JSON, got none")
	}
}
