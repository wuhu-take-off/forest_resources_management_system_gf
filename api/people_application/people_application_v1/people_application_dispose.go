package people_application_v1

import "github.com/gogf/gf/v2/frame/g"

type PeoplesApplicationDisposeReq struct {
	g.Meta                `path:"/people_application/dispose" method:"post" tags:"people_application"`
	PeopleApplicationId   int    `json:"peopleApplicationId"                  ` // 民事申请ID
	PeopleApplicationDesc string `json:"peopleApplicationDesc"      `           // 处理描述
}
type PeoplesApplicationDisposeRes struct{}
