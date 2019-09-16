package utils

import (
	"github.com/astaxie/beego"
	"os"
	"path/filepath"
)

/**
 * 重置beego相关配置参数<BR>
 * @author  作者: Jack Chu E-mail: chuchuanbao@gmail.com
 * 创建时间：2019/9/11 18:15
 */
func ResetConfig() {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var filename = "app.conf"
	if os.Getenv("BEEGO_RUNMODE") != "" {
		filename = os.Getenv("BEEGO_RUNMODE") + ".app.conf"
	}
	appConfigPath := filepath.Join(workPath, "conf", filename)
	beego.LoadAppConfig("ini", appConfigPath)
}
