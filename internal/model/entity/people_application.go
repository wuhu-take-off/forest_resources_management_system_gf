// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// PeopleApplication is the golang structure for table people_application.
type PeopleApplication struct {
	PeopleApplicationId      int    `json:"peopleApplicationId"      orm:"people_application_id"      ` // 民事申请ID
	PeopleApplicationTitle   string `json:"peopleApplicationTitle"   orm:"people_application_title"   ` // 民事申请标题
	PeopleApplicant          string `json:"peopleApplicant"          orm:"people_applicant"           ` // 申请人
	PeopleApplicationPhone   string `json:"peopleApplicationPhone"   orm:"people_application_phone"   ` // 联系电话
	PeopleApplicationType    int    `json:"peopleApplicationType"    orm:"people_application_type"    ` // 申请类型
	PeopleApplicationContent string `json:"peopleApplicationContent" orm:"people_application_content" ` // 申请内容
	PeopleApplicationAnnex   string `json:"peopleApplicationAnnex"   orm:"people_application_annex"   ` // 附件
	ModifyTime               int64  `json:"modifyTime"               orm:"modify_time"                ` // 修改时间
	PeopleApplicationState   int    `json:"peopleApplicationState"   orm:"people_application_state"   ` // 申请状态(0:待审核 1:已处理)
	PeopleApplicationDesc    string `json:"peopleApplicationDesc"    orm:"people_application_desc"    ` // 处理描述
}
