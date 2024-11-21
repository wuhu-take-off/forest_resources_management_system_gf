package identity_v1

import "github.com/gogf/gf/v2/frame/g"

type IdentityListReq struct {
	g.Meta `path:"/list" method:"post"  tags:"identity"`
}
type IdentityListStruct struct {
	IdentityId   int    `json:"identity_id"`
	IdentityName string `json:"identity_name"`
}
type IdentityListRes struct {
	IdentityList []*IdentityListStruct `json:"identity_list"`
}
