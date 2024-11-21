// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Notice is the golang structure for table notice.
type Notice struct {
	NoticeId        int    `json:"noticeId"        orm:"notice_id"        ` // 公告id
	NoticePublisher string `json:"noticePublisher" orm:"notice_publisher" ` // 发布人
	NoticeTime      int64  `json:"noticeTime"      orm:"notice_time"      ` // 发布日期
	NoticePriority  int    `json:"noticePriority"  orm:"notice_priority"  ` // 优先级
	NoticeContent   string `json:"noticeContent"   orm:"notice_content"   ` // 公告内容
	NoticeAnnex     string `json:"noticeAnnex"     orm:"notice_annex"     ` // 附件信息
}
