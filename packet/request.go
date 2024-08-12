package packet

import (
	"encoding/json"
	"errors"
)

type Request struct {
	Function string        `json:"function"`
	Args     []interface{} `json:"args"`
}

func NewRequest(function string, args []interface{}) (*Request, error) {
	if function == "" {
		return nil, errors.New("failed to new request:\n\tfunction name can't be empty")
	}
	return &Request{
		Function: function,
		Args:     args,
	}, nil
}

func (req *Request) MakeJson() ([]byte, error) {
	ret, err := json.Marshal(req)
	if err != nil {
		return []byte{}, errors.New("failed to make json:\n\t" + err.Error())
	}
	return ret, nil
}

func ParseRequest(msg []byte) (*Request, error) {
	var req Request
	err := json.Unmarshal(msg, &req)
	if err != nil {
		return nil, errors.New("failed to parse request:\n\t" + err.Error())
	}
	return &req, nil
}
