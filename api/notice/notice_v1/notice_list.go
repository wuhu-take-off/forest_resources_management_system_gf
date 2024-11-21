package notice_v1

import "github.com/gogf/gf/v2/frame/g"

type NoticeListReq struct {
	g.Meta `path:"/notice/list" method:"post" tags:"notice"`
	Page   int `json:"page"`
	Limit  int `json:"limit"`
}
type NoticeListRes struct {
	TotalCount int                 `json:"total_count"`
	NoticeList []*NoticeListStruct `json:"notice_list"`
}
type NoticeListStruct struct {
	NoticeId        int    `json:"noticeId"        ` // 公告ID
	NoticePublisher string `json:"noticePublisher" ` // 发布人
	NoticeTime      int64  `json:"noticeTime"      ` // 发布日期
	NoticePriority  int    `json:"noticePriority"  ` // 优先级
	NoticeContent   string `json:"noticeContent"   ` // 公告内容
	NoticeAnnex     string `json:"noticeAnnex"     ` // 附件信息
}
