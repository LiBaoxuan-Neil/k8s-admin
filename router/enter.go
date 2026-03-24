package router

import (
	"k8s-admin/router/example"
	"k8s-admin/router/k8s"
)

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
	K8sRouterGroup     k8s.K8sRouter
}

var RouterGroupApp = new(RouterGroup)
