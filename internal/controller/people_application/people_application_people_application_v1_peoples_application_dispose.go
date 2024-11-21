package people_application

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_people_application"

	"forest_resources_management_system_gf/api/people_application/people_application_v1"
)

func (c *ControllerPeople_application_v1) PeoplesApplicationDispose(ctx context.Context, req *people_application_v1.PeoplesApplicationDisposeReq) (res *people_application_v1.PeoplesApplicationDisposeRes, err error) {
	return l_people_application.NewPeopleApplicationLogic().PeoplesApplicationDispose(ctx, req)

}
