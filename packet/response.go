package packet

import (
	"encoding/json"
	"errors"
)

type Response struct {
	Condition string        `json:"condition"`
	Function  string        `json:"function"`
	Replys    []interface{} `json:"replys"`
	ErrorMsg  string        `json:"errormsg"`
}

func NewResponse(condition, function string, replys []interface{}, errormsg string) (*Response, error) {
	if condition == "" {
		return nil, errors.New("failed to new Response:\n\tcondition can't be empty")
	}
	if function == "" {
		return nil, errors.New("failed to new Response:\n\tfunction can't be empty")
	}
	return &Response{
		Condition: condition,
		Function:  function,
		Replys:    replys,
		ErrorMsg: errormsg,
	}, nil
}

func (res *Response) MakeJson() ([]byte, error) {
	ret, err := json.Marshal(res)
	if err != nil {
		return []byte{}, errors.New("failed to make json:\n\t" + err.Error())
	}
	return ret, nil
}
