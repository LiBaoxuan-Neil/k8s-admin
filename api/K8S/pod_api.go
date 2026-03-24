package k8s

import (
	"context"
	"fmt"
	"k8s-admin/global"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/gin-gonic/gin"
)

type PodApi struct {
}

func (*PodApi) GetPodList(c *gin.Context) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods("test").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, item := range list.Items {
		fmt.Println(item.Namespace, item.Name)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": list,
	})
}
