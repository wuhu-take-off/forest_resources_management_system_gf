package common

import "forest_resources_management_system_gf/internal/consts"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// NewResponse  code, msg, data/code Data
func NewResponse(code int, rest ...any) *Response {
	r := &Response{
		Code: code,
		Data: struct{}{},
	}
	if len(rest) > 0 {
		switch rest[0].(type) {
		case string:
			r.Msg = rest[0].(string)
		default:
			r.Data = rest[0]
		}
	} else if len(rest) > 1 {
		r.Msg = rest[0].(string)
		r.Data = rest[1]
	} else {
		r.Msg = consts.GetResponseMsg(code)
	}
	return r
}
