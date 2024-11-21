package people_application_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type PeopleApplicationCreateReq struct {
	g.Meta `path:"/people_application/create" method:"post" tags:"people_application" summary:"创建民事申请"`
	//PeopleApplicationId         int    `json:"peopleApplicationId"                  ` // 民事申请ID
	PeopleApplicationTitle   string              `json:"peopleApplicationTitle"            `         // 民事申请标题
	PeopleApplicant          string              `json:"peopleApplicant"                           ` // 申请人
	PeopleApplicationPhone   string              `json:"peopleApplicationPhone"            `         // 联系电话
	PeopleApplicationType    int                 `json:"peopleApplicationType"              `        // 申请类型
	PeopleApplicationContent string              `json:"peopleApplicationContent"        `           // 申请内容
	PeopleApplicationAnnex   []*ghttp.UploadFile `json:"peopleApplicationAnnex"  `                   // 附件
}
type PeopleApplicationCreateRes struct {
}
