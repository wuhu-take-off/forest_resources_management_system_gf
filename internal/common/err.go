package common

import "forest_resources_management_system_gf/internal/consts"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
func NewError(code int, message ...string) *Error {
	e := new(Error)
	e.Code = code
	if len(message) > 0 {
		e.Message = message[0]
	} else {
		e.Message = consts.GetResponseMsg(code)
	}
	return e
}
