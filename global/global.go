package global

import (
	"k8s-admin/config"

	"k8s.io/client-go/kubernetes"
)

var (
	CONF          *config.Server
	KubeConfigSet *kubernetes.Clientset
)
