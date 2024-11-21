package l_user

import (
	"context"
	"forest_resources_management_system_gf/api/user/user_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func (d *defaultUserV1) UserLogin(ctx context.Context, req *user_v1.UserLoginReq) (res *user_v1.UserLoginRes, err error) {
	var userData *do.User
	common.CopyFields(req, &userData)
	if one, err := dao.User.Ctx(ctx).One(userData); err != nil {
		g.Log().Errorf(ctx, "用户登录失败: %v", err)
		return nil, common.NewError(consts.InvalidData, "用户名或密码错误")
	} else if one == nil {
		return nil, common.NewError(consts.LoginFailed, "用户名或密码错误")
	} else {
		var userInfo *entity.User
		one.Struct(&userInfo)
		token, err := common.GenerateToken(userInfo.UserId, userInfo.UserName)
		if err != nil {
			return nil, err
		}
		return &user_v1.UserLoginRes{
			UserId:   userInfo.UserId,
			UserName: userInfo.UserName,
			Token:    token,
		}, nil
	}
}
