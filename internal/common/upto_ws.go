package common

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	"net/http"
)

func UptoWS(r *ghttp.Request) (*websocket.Conn, error) {
	//0. 初始化 websocket 配置
	upgrade := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		//检查origin 限定websocket 被其他的域名访问
		return true
	}}
	//1. 升级连接为 websocket
	return upgrade.Upgrade(r.Response.Writer, r.Request, http.Header{
		//获取前端请求头的Sec-Websocket-Protocol值
		"Sec-Websocket-Protocol": {r.Request.Header.Get("Sec-Websocket-Protocol")},
	})
}
