package main

import (
	"fmt"
	"os"
	"path/filepath"

	"corpweb/conf"
	_ "corpweb/routers"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

// 加载项目根目录
func loadProjectHome() {
	wd, err := os.Getwd()
	if err != nil {
		beego.Error(err)
	}
	conf.ProjectHome = filepath.Clean(wd)
	beego.Info(fmt.Sprintf("conf(ProjectHome): %s", conf.ProjectHome))
}

func main() {
	loadProjectHome()
	conf.LoadAllConf()

	beego.SetLogger("file", `{"filename":"app.log"}`)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)

	beego.SetStaticPath("/res", conf.ResourceHome)
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
