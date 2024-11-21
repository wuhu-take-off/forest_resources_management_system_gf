package notice_v1

import "github.com/gogf/gf/v2/frame/g"

type NoticeInfoListReq struct {
	g.Meta         `path:"/notice/info/list" method:"post" tags:"notice" `
	Page           int    `json:"page"`            // 页码
	Limit          int    `json:"limit"`           // 每页数量
	NoticePriority *int   `json:"notice_priority"` // 优先级
	ExpirationDate *int64 `json:"expiration_date"` // 过期时间
}
type NoticeInfoListStruct struct {
	NoticeId        int    `json:"noticeId"        ` // 公告ID
	NoticePublisher string `json:"noticePublisher" ` // 发布人
	NoticeTime      int64  `json:"noticeTime"      ` // 发布日期
	NoticePriority  int    `json:"noticePriority"  ` // 优先级
	NoticeContent   string `json:"noticeContent"   ` // 公告内容
	NoticeAnnex     string `json:"noticeAnnex"     ` // 附件信息
}
type NoticeInfoListRes struct {
	TotalCount     int                     `json:"total_count"`
	NoticeInfoList []*NoticeInfoListStruct `json:"notice_info_list"`
}
