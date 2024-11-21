package middleware

import (
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	TokenClaims = "token_claims"
)

var skipPaths = map[string]struct{}{
	"/chat": {},
}

func TokenAuthMiddleware(r *ghttp.Request) {
	//跳过不需要验证的路径
	if _, ok := skipPaths[r.Request.URL.Path]; ok {
		println("skip token auth for path:", r.Request.URL.Path)
		r.Middleware.Next()
		return
	}

	//获取token
	token := r.Header.Get("Authorization")
	if token == "" {
		r.SetError(common.NewError(consts.Unauthorized))
		sendResponse(r)
		return
	} else {
		if userInfo, err := common.ParseToken(token); err != nil {
			r.SetError(common.NewError(consts.Unauthorized))
			sendResponse(r)
			return
		} else {
			r.SetCtxVar(TokenClaims, userInfo)
		}
	}
	r.Middleware.Next()
}
