package l_people_application

import (
	"context"
	"forest_resources_management_system_gf/api/people_application/people_application_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultPeopleApplication) PeopleApplicationDel(ctx context.Context, req *people_application_v1.PeopleApplicationDelReq) (res *people_application_v1.PeopleApplicationDelRes, err error) {
	var peopleApplication *do.PeopleApplication
	common.CopyFields(req, &peopleApplication)
	if _, err = dao.PeopleApplication.Ctx(ctx).Delete(peopleApplication); err != nil {
		g.Log().Errorf(ctx, "delete PeopleApplication info failed, err:%v", err)
		return nil, common.NewError(consts.InternalServerError, "delete PeopleApplication info failed")
	}
	return
}
