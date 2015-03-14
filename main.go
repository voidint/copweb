package main

import (
	"corpweb/conf"

	_ "git.oschina.net/voidint/cms4go/routers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func main() {
	beego.SetLogger("file", `{"filename":"app.log"}`)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)

	beego.SetStaticPath("/res", conf.Resource_Home)
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
