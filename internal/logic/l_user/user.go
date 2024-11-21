package l_user

import (
	"forest_resources_management_system_gf/api/user"
)

type defaultUserV1 struct{}
type defaultUserV2 struct{}

type IUserLogicV1 interface {
	user.IUserUser_v1
	//UserLogin(ctx context.Context, req *user_v1.UserLoginReq) (res *user_v1.UserLoginRes, err error)
}
type IUserLogicV2 interface {
	user.IUserUser_v2
	//UserCreate(ctx context.Context, req *user_v2.UserCreateReq) (res *user_v2.UserCreateRes, err error)
	//UserDel(ctx context.Context, req *user_v2.UserDelReq) (res *user_v2.UserDelRes, err error)
	//UserList(ctx context.Context, req *user_v2.UserListReq) (res *user_v2.UserListRes, err error)
	//UserModify(ctx context.Context, req *user_v2.UserModifyReq) (res *user_v2.UserModifyRes, err error)
}

func NewUserLogicV1() IUserLogicV1 {
	return &defaultUserV1{}
}
func NewUserLogicV2() IUserLogicV2 {
	return &defaultUserV2{}
}
