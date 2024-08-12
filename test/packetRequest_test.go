package test

import (
	"github.com/Alonza0314/lotus/packet"
	"testing"
)

func TestNewRequest(t *testing.T) {
	function := "testFunction"
	parameters := []interface{}{1, "param"}
	req, err := packet.NewRequest(function, parameters)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if req.Function != function {
		t.Errorf("expected function %v, got %v", function, req.Function)
	}
	if len(req.Parameters) != len(parameters) {
		t.Errorf("expected parameters length %v, got %v", len(parameters), len(req.Parameters))
	}

	_, err = packet.NewRequest("", parameters)
	if err == nil {
		t.Fatal("expected error for empty function name, got none")
	}
}

func TestParseRequest(t *testing.T) {
	msg := `{"function":"testFunction","parameters":[1,"param"]}`
	expectedFunction := "testFunction"
	expectedParameters := []interface{}{1.0, "param"}

	req, err := packet.ParseRequest([]byte(msg))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if req.Function != expectedFunction {
		t.Errorf("expected function %v, got %v", expectedFunction, req.Function)
	}
	if len(req.Parameters) != len(expectedParameters) {
		t.Errorf("expected parameters length %v, got %v", len(expectedParameters), len(req.Parameters))
	}
	for i, param := range req.Parameters {
		if param != expectedParameters[i] {
			t.Errorf("expected parameter %v at index %d, got %v", expectedParameters[i], i, param)
		}
	}

	invalidMsg := `{"function":"testFunction","parameters":[1,"param"]`
	_, err = packet.ParseRequest([]byte(invalidMsg))
	if err == nil {
		t.Fatal("expected error for invalid JSON, got none")
	}
}
