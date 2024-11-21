package middleware

import (
	"errors"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ResponseMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
	sendResponse(r)
}
func sendResponse(r *ghttp.Request) {
	response := common.NewResponse(consts.Success)
	defer func() {
		r.Response.ClearBuffer()           // clear buffer
		r.Response.WriteJsonExit(response) // write response and exit
	}()

	errTmp := r.GetError()
	if errTmp == nil {
		response.Data = r.GetHandlerResponse()
		return
	}
	err := new(common.Error)
	if !errors.As(errTmp, &err) {
		response = common.NewResponse(consts.InternalServerError)
		return
	} else {
		response.Code = err.Code
		response.Msg = err.Message
	}
}
