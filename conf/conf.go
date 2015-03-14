package conf

import (
	"os"

	"github.com/astaxie/beego"
)

var (
	Resource_Home string = "resource"
)

func init() {
	LoadAllConf()
}

func LoadAllConf() {
	LoadResourceHome()
}

func LoadResourceHome() {
	if res := os.Getenv("RESOURCE_HOME"); res != "" {
		Resource_Home = res
	}
	beego.Info("RESOURCE_HOME = " + Resource_Home)
}
