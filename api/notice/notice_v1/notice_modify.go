package notice_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type NoticeModifyReq struct {
	g.Meta          `path:"/notice/modify" method:"post" tags:"notice" `
	NoticeId        int                 `json:"noticeId"        ` // 公告ID
	NoticePublisher string              `json:"noticePublisher" ` // 发布人
	NoticeTime      int64               `json:"noticeTime"      ` // 发布日期
	NoticePriority  int                 `json:"noticePriority"  ` // 优先级
	NoticeContent   string              `json:"noticeContent"   ` // 公告内容
	NoticeAnnex     []*ghttp.UploadFile `json:"noticeAnnex"     ` // 附件信息
}
type NoticeModifyRes struct{}
