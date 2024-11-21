package l_people_application

import (
	"forest_resources_management_system_gf/api/people_application"
)

type defaultPeopleApplication struct{}

type IPeopleApplicationLogic interface {
	people_application.IPeopleApplicationPeople_application_v1
}

func NewPeopleApplicationLogic() IPeopleApplicationLogic {
	return &defaultPeopleApplication{}
}
