package char_v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CallChatReq struct {
	g.Meta `path:"/chat" method:"get" tags:"chat"`
}
type CallChatRes struct{}

type ChatContent struct {
	Type         int    `json:"type"` // 1:text 2:file
	SendId       int    `json:"send_id"`
	DepartmentId *int   `json:"department_id"`
	ReceiverId   *int   `json:"receiver_id"`
	SendTime     int64  `json:"send_time"`
	Content      string `json:"msg"` // text or file content
}
type ChatMessage struct {
	Type    int    `json:"type"`    // 1:text 2:file
	Content string `json:"content"` // text content or file url
}
