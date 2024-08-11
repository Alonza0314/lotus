package packet

import (
	"encoding/json"
	"errors"
)

type Request struct {
	Function   string        `json:"function"`
	Parameters []interface{} `json:"parameters"`
}

func NewRequest(function string, parameters []interface{}) (*Request, error) {
	if function == "" {
		return nil, errors.New("failed to new request:\n\tfunction name can't be empty")
	}
	return &Request{
		Function: function,
		Parameters: parameters,
	}, nil
}

func ParseRequest(msg []byte) (*Request, error) {
	var req Request
	err := json.Unmarshal(msg, &req)
	if err != nil {
		return nil, errors.New("failed to parse request:\n\t" + err.Error())
	}
	return &req, nil
}