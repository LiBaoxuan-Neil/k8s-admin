package initiallize

import (
	"k8s-admin/config"
	"k8s-admin/global"

	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	global.CONF = &config.Server{}
	err = v.Unmarshal(global.CONF)
	if err != nil {
		panic(err)
	}
	return v
}
