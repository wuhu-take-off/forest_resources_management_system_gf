package l_people_application

import (
	"context"
	"forest_resources_management_system_gf/api/people_application/people_application_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultPeopleApplication) PeopleApplicationList(ctx context.Context, req *people_application_v1.PeopleApplicationListReq) (res *people_application_v1.PeopleApplicationListRes, err error) {
	var peopleApplications []*entity.PeopleApplication
	db := dao.PeopleApplication.Ctx(ctx)
	if req.PeopleApplicationState != nil {
		db = db.Where(dao.PeopleApplication.Columns().PeopleApplicationState, *req.PeopleApplicationState)
	}

	if err := db.Page(req.Page, req.Limit).
		Order(dao.PeopleApplication.Columns().PeopleApplicationId).Scan(&peopleApplications); err != nil {
		g.Log().Errorf(ctx, "PeopleApplicationList Error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "PeopleApplicationList Error")
	}
	res = &people_application_v1.PeopleApplicationListRes{
		PeopleApplicationList: make([]*people_application_v1.PeopleApplicationStruct, 0),
	}
	common.CopyFields(peopleApplications, &res.PeopleApplicationList)

	if total, err := db.Count(); err != nil {
		g.Log().Errorf(ctx, "PeopleApplicationList  Error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "PeopleApplicationList Error")
	} else {
		res.TotalCount = total
	}

	//policyInfo.PolicyTime = common.DateConversion(req.PolicyTime).(int64)
	return
}
