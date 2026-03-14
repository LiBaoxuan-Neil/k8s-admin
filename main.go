package main

import (
	"k8s-admin/global"
	"k8s-admin/initiallize"
)

func main() {
	initiallize.Viper()
	initiallize.K8S()
	r := initiallize.Routers(EmbedFrontend)
	panic(r.Run(global.CONF.System.Addr))
}
