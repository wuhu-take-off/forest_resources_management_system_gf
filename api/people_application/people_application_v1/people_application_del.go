package people_application_v1

import "github.com/gogf/gf/v2/frame/g"

type PeopleApplicationDelReq struct {
	g.Meta              `path:"/people_application/del" method:"post" tags:"people_application"`
	PeopleApplicationId int `json:"peopleApplicationId"                ` // 民事申请ID
}
type PeopleApplicationDelRes struct {
}
