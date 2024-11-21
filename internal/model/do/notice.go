// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Notice is the golang structure of table notice for DAO operations like Where/Data.
type Notice struct {
	g.Meta          `orm:"table:notice, do:true"`
	NoticeId        interface{} // 公告id
	NoticePublisher interface{} // 发布人
	NoticeTime      interface{} // 发布日期
	NoticePriority  interface{} // 优先级
	NoticeContent   interface{} // 公告内容
	NoticeAnnex     interface{} // 附件信息
}
