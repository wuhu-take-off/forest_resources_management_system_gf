package people_application_v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PeopleApplicationStruct struct {
	PeopleApplicationId      int    `json:"peopleApplicationId"                  `      // 民事申请ID
	PeopleApplicationTitle   string `json:"peopleApplicationTitle"            `         // 民事申请标题
	PeopleApplicant          string `json:"peopleApplicant"                           ` // 申请人
	PeopleApplicationPhone   string `json:"peopleApplicationPhone"            `         // 联系电话
	PeopleApplicationType    int    `json:"peopleApplicationType"              `        // 申请类型
	PeopleApplicationContent string `json:"peopleApplicationContent"        `           // 申请内容
	PeopleApplicationAnnex   string `json:"peopleApplicationAnnex"  `                   // 附件
	ModifyTime               int64  `json:"modifyTime"                           `      // 修改时间
	PeopleApplicationState   int    `json:"peopleApplicationState"    `                 // 申请状态(0:待审核 1:已处理)
	PeopleApplicationDesc    string `json:"peopleApplicationDesc"      `                // 处理描述
}

type PeopleApplicationListReq struct {
	g.Meta                 `path:"/people_application/list" method:"post"  tags:"people_application"`
	Page                   int  `json:"page" `
	Limit                  int  `json:"limit" `
	PeopleApplicationState *int `json:"people_application_state" ` // 申请状态(0:待审核 1:已处理)
}
type PeopleApplicationListRes struct {
	TotalCount            int                        `json:"total_count"`
	PeopleApplicationList []*PeopleApplicationStruct `json:"people_application_list"`
}
