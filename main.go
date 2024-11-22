package main

import (
	"forest_resources_management_system_gf/internal/controller/chat"
	"forest_resources_management_system_gf/internal/controller/department"
	"forest_resources_management_system_gf/internal/controller/forest"
	"forest_resources_management_system_gf/internal/controller/forest_resource_info"
	"forest_resources_management_system_gf/internal/controller/identity"
	"forest_resources_management_system_gf/internal/controller/notice"
	"forest_resources_management_system_gf/internal/controller/people_application"
	"forest_resources_management_system_gf/internal/controller/policy"
	"forest_resources_management_system_gf/internal/controller/resource"
	"forest_resources_management_system_gf/internal/controller/tree_species"
	"forest_resources_management_system_gf/internal/controller/user"
	"forest_resources_management_system_gf/middleware"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func main() {

	// 加载配置文件
	s := g.Server()
	//设置跨域
	s.BindMiddlewareDefault(MiddlewareCORS)
	s.Group("/user", func(group *ghttp.RouterGroup) {
		group.Bind(user.NewUser_v1()).Middleware(middleware.ResponseMiddleware)
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.ResponseMiddleware)
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.Bind(user.NewUser_v2()).Middleware(middleware.TokenAuthMiddleware)
		})
		group.Group("/department", func(group *ghttp.RouterGroup) {
			group.Bind(department.NewDepartment_v1())
		})
		group.Group("/identity", func(group *ghttp.RouterGroup) {
			group.Bind(identity.NewIdentity_v1())
		})
		group.Bind(forest.NewForest_v1())
		group.Bind(tree_species.NewTree_species_v1())
		group.Bind(forest_resource_info.NewForest_resource_info_v1())
		group.Bind(resource.NewResource_v1())
		group.Bind(policy.NewPolicy_v1())
		group.Bind(notice.NewNotice_v1())
		group.Bind(people_application.NewPeople_application_v1())
		group.Bind(chat.NewChar_v1()).Middleware(middleware.TokenAuthMiddleware)
	})
	s.Run()
}
