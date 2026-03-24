package k8s

import (
	"context"
	"k8s-admin/global"
	"k8s-admin/response"

	namespace_res "k8s-admin/model/namespace/response"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceAPi struct {
}

func (*NamespaceAPi) GetNamespaceList(c *gin.Context) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	namespaceList := make([]namespace_res.NameSpace, 0)
	for _, item := range list.Items {
		namespaceList = append(namespaceList, namespace_res.NameSpace{
			Name:              item.Name,
			CreationTimestamp: item.CreationTimestamp.Unix(),
			Status:            string(item.Status.Phase),
		})
	}
	response.SuccessWithData(c, namespaceList)
}
