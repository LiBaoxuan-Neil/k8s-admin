package router

import "k8s-admin/router/example"

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
}

var RouterGroupApp = new(RouterGroup)
