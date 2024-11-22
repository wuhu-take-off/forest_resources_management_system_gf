package l_tst

import (
	"forest_resources_management_system_gf/api/tst"
)

type defaultTst struct{}

type ITstLogic interface {
	tst.ITstTst_v1
}

func NewTstLogic() ITstLogic {
	return &defaultTst{}
}
