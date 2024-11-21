package l_notice

import (
	"context"
	"encoding/json"
	"forest_resources_management_system_gf/api/notice/notice_v1"
	"forest_resources_management_system_gf/api/universal_class/universal_class_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
	"time"
)

func (d defaultNotice) NoticeModify(ctx context.Context, req *notice_v1.NoticeModifyReq) (res *notice_v1.NoticeModifyRes, err error) {
	t := strconv.FormatInt(time.Now().Unix(), 10) + "/"
	annex := make([]*universal_class_v1.Annex, 0)

	// 保存附件
	for i := range req.NoticeAnnex {
		if _, err := req.NoticeAnnex[i].Save(consts.NoticeAnnexDir + t); err != nil {
			g.Log().Errorf(ctx, "Failed to save notice annex: %v", err)
			return nil, common.NewError(consts.InternalServerError, "Failed to save notice annex")
		}
		annex = append(annex, &universal_class_v1.Annex{
			Url:  consts.NoticeAnnexDir + t,
			Name: req.NoticeAnnex[i].Filename,
		})
	}
	// 保存资源信息
	var noticeInfo *do.Notice
	common.CopyFields(req, &noticeInfo)

	if len(annex) > 0 {
		annexJson, _ := json.Marshal(annex)
		noticeInfo.NoticeAnnex = string(annexJson)
	}
	noticeInfo.NoticeId = nil

	if _, err = dao.Notice.Ctx(ctx).Where(dao.Notice.Columns().NoticeId, req.NoticeId).Update(noticeInfo); err != nil {
		g.Log().Errorf(ctx, "Failed to create notice info: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Failed to create notice info")
	}
	return
}
