package l_policy

import (
	"forest_resources_management_system_gf/api/policy"
)

type defaultPolicy struct{}

type IPolicyLogic interface {
	policy.IPolicyPolicy_v1
}

func NewPolicyLogic() IPolicyLogic {
	return &defaultPolicy{}
}
