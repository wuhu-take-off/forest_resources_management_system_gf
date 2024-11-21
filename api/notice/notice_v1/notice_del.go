package notice_v1

import "github.com/gogf/gf/v2/frame/g"

type NoticeDelReq struct {
	g.Meta   `path:"/notice/del" method:"post" tags:"notice" `
	NoticeId int `json:"notice_id" `
}
type NoticeDelRes struct {
}
