package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
)

const (
	// 默认的资源根目录为项目下的resource目录
	DefaultResourceHome = "resource"
	// 资源根目录环境变量名称
	EnvResourceHome = "RESOURCE_HOME"
)

var (
	// 项目根目录
	ProjectHome string
	// 系统资源目录
	ResourceHome string
)

// 加载所有除conf/app.conf文件之外的配置信息
func LoadAllConf() {
	LoadResourceHome()
}

// 加载系统资源目录路径
func LoadResourceHome() {
	if res := os.Getenv(EnvResourceHome); len(res) > 0 {
		ResourceHome = filepath.Clean(res)
	} else {
		ResourceHome = filepath.Clean(ProjectHome) + string(os.PathSeparator) + DefaultResourceHome
	}

	beego.Info(fmt.Sprintf("conf(ResourceHome): %s", ResourceHome))

	// 创建必要的资源目录
	srcImgDir := ResourceHome + string(os.PathSeparator) + "image" + string(os.PathSeparator) + "source"
	err := os.MkdirAll(srcImgDir, 0755)
	if err != nil {
		beego.Error(fmt.Sprintf("os.MkdirAll(%s, 0755) err:%s", srcImgDir, err))
	}
}
