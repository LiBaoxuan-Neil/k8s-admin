package k8s

import (
	"k8s-admin/api"

	"github.com/gin-gonic/gin"
)

type K8sRouter struct {
}

func (*K8sRouter) Init(r *gin.RouterGroup) {
	apiRouter := r.Group("/k8s")
	apiGroup := api.ApiGroupApp.K8sApiGroup
	// 命名空间列表接口
	// 1. Pod创建
	// 2. Pod查看（详情/列表）
	// 3. Pod编辑（更新/升级）
	// 4. Pod删除
	apiRouter.GET("/pod", apiGroup.GetPodList)
	apiRouter.GET("namespaces", apiGroup.GetNamespaceList)
}
