package department_affiche_v1

type DepartmentAfficheListReq struct{}

type DepartmentAfficheListStruct struct {
	DepartmentAfficheId      int    `json:"department_affiche_id"         ` // 部门公告主键
	DepartmentAfficheTitle   string `json:"department_affiche_title"      ` // 部门公告标题
	DepartmentAfficheContent string `json:"department_affiche_content"    ` // 部门公告内容
	DepartmentId             int    `json:"department_id"                `  // 发布部门
	//UserId                      int    `json:""                      ` // 创建用户
	DepartmentAfficheCreateTime int64 `json:"department_affiche_create_time" ` // 创建时间
}

type DepartmentAfficheListRes struct {
	DepartmentAfficheList []*DepartmentAfficheListStruct `json:"department_affiche_list"`
}
