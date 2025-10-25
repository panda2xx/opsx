package main

import (
	"os"

	"github.com/panda2xx/opsx/cmd/opsx-center/app"
)

func main() {
	cmd := app.NewOpsXCenterCommand()

	if err := cmd.Execute(); err != nil {
		// 如果发生错误，则退出程序
		// 返回退出码，可以使其他程序（例如 bash 脚本）根据退出码来判断服务运行状态
		os.Exit(1)
	}
}
