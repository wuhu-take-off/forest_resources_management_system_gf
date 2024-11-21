package people_application

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_people_application"

	"forest_resources_management_system_gf/api/people_application/people_application_v1"
)

func (c *ControllerPeople_application_v1) PeopleApplicationCreate(ctx context.Context, req *people_application_v1.PeopleApplicationCreateReq) (res *people_application_v1.PeopleApplicationCreateRes, err error) {
	return l_people_application.NewPeopleApplicationLogic().PeopleApplicationCreate(ctx, req)
}
