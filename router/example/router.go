package example

import (
	"k8s-admin/api"

	"github.com/gin-gonic/gin"
)

type ExampleRouter struct {
}

func (s *ExampleRouter) InitExampleRouter(r *gin.RouterGroup) {
	apiRouter := r.Group("/example")
	apiGroup := api.ApiGroupApp.ExampleApiGroup
	apiRouter.GET("/ping", apiGroup.ExampleTest)
}
