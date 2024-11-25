// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"forest_resources_management_system_gf/api/user/user_v1"
	"forest_resources_management_system_gf/api/user/user_v2"
)

type IUserUser_v1 interface {
	UserLogin(ctx context.Context, req *user_v1.UserLoginReq) (res *user_v1.UserLoginRes, err error)
}

type IUserUser_v2 interface {
	SingleUserInfo(ctx context.Context, req *user_v2.SingleUserInfoReq) (res *user_v2.SingleUserInfoRes, err error)
	UserChatList(ctx context.Context, req *user_v2.UserChatListReq) (res *user_v2.UserChatListRes, err error)
	UserCreate(ctx context.Context, req *user_v2.UserCreateReq) (res *user_v2.UserCreateRes, err error)
	UserDel(ctx context.Context, req *user_v2.UserDelReq) (res *user_v2.UserDelRes, err error)
	UserList(ctx context.Context, req *user_v2.UserListReq) (res *user_v2.UserListRes, err error)
	UserModify(ctx context.Context, req *user_v2.UserModifyReq) (res *user_v2.UserModifyRes, err error)
}
