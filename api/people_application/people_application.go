// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package people_application

import (
	"context"

	"forest_resources_management_system_gf/api/people_application/people_application_v1"
)

type IPeopleApplicationPeople_application_v1 interface {
	PeopleApplicationCreate(ctx context.Context, req *people_application_v1.PeopleApplicationCreateReq) (res *people_application_v1.PeopleApplicationCreateRes, err error)
	PeopleApplicationDel(ctx context.Context, req *people_application_v1.PeopleApplicationDelReq) (res *people_application_v1.PeopleApplicationDelRes, err error)
	PeoplesApplicationDispose(ctx context.Context, req *people_application_v1.PeoplesApplicationDisposeReq) (res *people_application_v1.PeoplesApplicationDisposeRes, err error)
	PeopleApplicationList(ctx context.Context, req *people_application_v1.PeopleApplicationListReq) (res *people_application_v1.PeopleApplicationListRes, err error)
	PeopleApplicationModify(ctx context.Context, req *people_application_v1.PeopleApplicationModifyReq) (res *people_application_v1.PeopleApplicationModifyRes, err error)
}
