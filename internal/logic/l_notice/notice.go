package l_notice

import (
	"forest_resources_management_system_gf/api/notice"
)

type defaultNotice struct{}

type INoticeLogic interface {
	notice.INoticeNotice_v1
}

func NewNoticeLogic() INoticeLogic {
	return &defaultNotice{}
}
