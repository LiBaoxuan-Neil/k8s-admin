package initiallize

import (
	"io/fs"

	"k8s-admin/middleware"
	"k8s-admin/router"

	"github.com/gin-gonic/gin"
)

func Routers(frontendFS fs.FS) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors)

	// API 挂到 /api 下，与前端 VUE_APP_BASE_API='/api' 一致
	apiGroup := r.Group("/api")
	exampleRouterGroup := router.RouterGroupApp.ExampleRouterGroup
	exampleRouterGroup.InitExampleRouter(apiGroup)
	k8sRouterGroup := router.RouterGroupApp.K8sRouterGroup
	k8sRouterGroup.Init(apiGroup)
	// 嵌入的前端静态资源 + SPA 回退
	MountFrontend(r, frontendFS)

	return r
}
