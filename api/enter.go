package api

import (
	k8s "k8s-admin/api/K8S"
	"k8s-admin/api/example"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8sApiGroup     k8s.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
