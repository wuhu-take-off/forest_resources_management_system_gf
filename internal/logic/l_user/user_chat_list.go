package l_user

import (
	"context"
	"forest_resources_management_system_gf/api/user/user_v2"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"forest_resources_management_system_gf/middleware"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultUserV2) UserChatList(ctx context.Context, req *user_v2.UserChatListReq) (res *user_v2.UserChatListRes, err error) {
	claims := ctx.Value(middleware.TokenClaims).(*common.Claims)
	if claims == nil {
		return nil, common.NewError(consts.Unauthorized, "token is invalid")
	}
	var userIdList []int

	// Fetch public chat list
	if array, err := dao.PrivateChat.Ctx(ctx).Where(dao.PrivateChat.Columns().PrivateChatSendId, claims.UserId).
		//WhereOr(dao.PrivateChat.Columns().PrivateChatReceiverId, claims.UserId).Distinct().
		Fields(dao.PrivateChat.Columns().PrivateChatReceiverId).Distinct().Array(); err != nil {
		g.Log().Errorf(ctx, "Error while fetching private chat list: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Error while fetching private chat list")
	} else {
		for i := range array {
			userIdList = append(userIdList, array[i].Int())
		}
	}
	// Fetch private chat list
	if array, err := dao.PrivateChat.Ctx(ctx).Where(dao.PrivateChat.Columns().PrivateChatReceiverId, claims.UserId).
		Fields(dao.PrivateChat.Columns().PrivateChatSendId).Distinct().Array(); err != nil {
		g.Log().Errorf(ctx, "Error while fetching private chat list: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Error while fetching private chat list")
	} else {
		for i := range array {
			userIdList = append(userIdList, array[i].Int())
		}
	}
	//数组去重
	distinctMap := make(map[int]struct{})
	for i := range userIdList {
		distinctMap[userIdList[i]] = struct{}{}
	}
	userIdList = make([]int, 0, len(distinctMap))
	for k := range distinctMap {
		userIdList = append(userIdList, k)
	}
	//获取用户名
	var userListTmp []*entity.User
	if err := dao.User.Ctx(ctx).WhereIn(dao.User.Columns().UserId, userIdList).
		Fields(dao.User.Columns().UserId, dao.User.Columns().UserName).Scan(&userListTmp); err != nil {
		g.Log().Errorf(ctx, "Error while fetching user list: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Error while fetching user list")
	}
	res = &user_v2.UserChatListRes{
		UserChatList: make([]*user_v2.UserChatListStruct, 0),
	}
	common.CopyFields(userListTmp, &res.UserChatList)
	return
}
