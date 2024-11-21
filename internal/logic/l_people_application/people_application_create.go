package l_people_application

import (
	"context"
	"encoding/json"
	"forest_resources_management_system_gf/api/people_application/people_application_v1"
	"forest_resources_management_system_gf/api/universal_class/universal_class_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
	"time"
)

func (d defaultPeopleApplication) PeopleApplicationCreate(ctx context.Context, req *people_application_v1.PeopleApplicationCreateReq) (res *people_application_v1.PeopleApplicationCreateRes, err error) {
	t := strconv.FormatInt(time.Now().Unix(), 10) + "/"
	annex := make([]*universal_class_v1.Annex, 0)

	// 保存附件
	for i := range req.PeopleApplicationAnnex {
		if _, err := req.PeopleApplicationAnnex[i].Save(consts.PeopleApplicationDir + t); err != nil {
			g.Log().Errorf(ctx, "Failed to save people applicant annex: %v", err)
			return nil, common.NewError(consts.InternalServerError, "Failed to save people applicant annex")
		}
		annex = append(annex, &universal_class_v1.Annex{
			Url:  consts.PeopleApplicationDir + t,
			Name: req.PeopleApplicationAnnex[i].Filename,
		})
	}
	// 保存资源信息
	var peopleApplication *do.PeopleApplication
	common.CopyFields(req, &peopleApplication)
	annexJson, _ := json.Marshal(annex)
	peopleApplication.PeopleApplicationAnnex = string(annexJson)
	peopleApplication.PeopleApplicationId = nil
	peopleApplication.ModifyTime = time.Now().Unix()
	peopleApplication.PeopleApplicationDesc = ""

	if _, err = dao.PeopleApplication.Ctx(ctx).Insert(peopleApplication); err != nil {
		g.Log().Errorf(ctx, "Failed to create people application info: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Failed to create people application info")
	}
	return
}