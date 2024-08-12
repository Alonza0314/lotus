package test

import (
	"encoding/json"
	"github.com/Alonza0314/lotus/packet"
	"testing"
)

func TestNewResponse(t *testing.T) {
	condition := "success"
	function := "testFunction"
	replys := []interface{}{1, "reply"}
	errormsg := "no error"
	res, err := packet.NewResponse(condition, function, replys, errormsg)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.Condition != condition {
		t.Errorf("expected condition %v, got %v", condition, res.Condition)
	}
	if res.Function != function {
		t.Errorf("expected function %v, got %v", function, res.Function)
	}
	if len(res.Replys) != len(replys) {
		t.Errorf("expected replys length %v, got %v", len(replys), len(res.Replys))
	}
	for i, reply := range res.Replys {
		if reply != replys[i] {
			t.Errorf("expected reply %v at index %d, got %v", replys[i], i, reply)
		}
	}
	if res.ErrorMsg != errormsg {
		t.Errorf("expected ErrorMsg %v, got %v", errormsg, res.ErrorMsg)
	}

	_, err = packet.NewResponse("", function, replys, errormsg)
	if err == nil {
		t.Fatal("expected error for empty condition, got none")
	}

	_, err = packet.NewResponse(condition, "", replys, errormsg)
	if err == nil {
		t.Fatal("expected error for empty function name, got none")
	}
}

func TestMakeJsonRes(t *testing.T) {
	condition := "success"
	function := "testFunction"
	replys := []interface{}{1, "reply"}
	errormsg := "no error"
	res, _ := packet.NewResponse(condition, function, replys, errormsg)
	jsonData, err := res.MakeJson()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var parsedResponse packet.Response
	if err := json.Unmarshal(jsonData, &parsedResponse); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	expectedReplys := []interface{}{1.0, "reply"}

	if parsedResponse.Condition != condition {
		t.Errorf("expected condition %v, got %v", condition, parsedResponse.Condition)
	}
	if parsedResponse.Function != function {
		t.Errorf("expected function %v, got %v", function, parsedResponse.Function)
	}
	if len(parsedResponse.Replys) != len(expectedReplys) {
		t.Errorf("expected replys length %v, got %v", len(expectedReplys), len(parsedResponse.Replys))
	}
	for i, reply := range parsedResponse.Replys {
		if reply != expectedReplys[i] {
			t.Errorf("expected reply %v at index %d, got %v", expectedReplys[i], i, reply)
		}
	}
	if parsedResponse.ErrorMsg != errormsg {
		t.Errorf("expected ErrorMsg %v, got %v", errormsg, parsedResponse.ErrorMsg)
	}
}
