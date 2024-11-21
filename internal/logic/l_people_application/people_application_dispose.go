package l_people_application

import (
	"context"
	"forest_resources_management_system_gf/api/people_application/people_application_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (d defaultPeopleApplication) PeoplesApplicationDispose(ctx context.Context, req *people_application_v1.PeoplesApplicationDisposeReq) (res *people_application_v1.PeoplesApplicationDisposeRes, err error) {

	// 保存资源信息
	var peopleApplication *do.PeopleApplication
	common.CopyFields(req, &peopleApplication)

	peopleApplication.PeopleApplicationId = nil
	peopleApplication.ModifyTime = time.Now().Unix()
	peopleApplication.PeopleApplicationState = 1

	if _, err = dao.PeopleApplication.Ctx(ctx).Where(dao.PeopleApplication.Columns().PeopleApplicationId, req.PeopleApplicationId).Update(peopleApplication); err != nil {
		g.Log().Errorf(ctx, "Failed to create people application info: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Failed to create people application info")
	}
	return
}
