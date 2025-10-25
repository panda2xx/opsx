package app

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/panda2xx/opsx/cmd/opsx-center/app/options"
)

var configFile string // 配置文件路径

// NewOpsXCenterCommand 创建一个 *cobra.Command 对象，用于启动应用程序.
func NewOpsXCenterCommand() *cobra.Command {
	// 创建默认的应用命令行选项
	opts := options.NewServerOptions()

	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "opsx-center",
		// 命令的简短描述
		Short: "devops center",
		Long:  `devops center, designed to manage devops resources`,
		// 命令出错时，不打印帮助信息。设置为 true 可以确保命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			// 将 viper 中的配置解析到选项 opts 变量中.
			if err := viper.Unmarshal(opts); err != nil {
				return err
			}

			// 对opts进行校验.
			if err := opts.Validate(); err != nil {
				return err
			}

			fmt.Printf("Read MySQL host from Viper: %s\n", viper.GetString("mysql.addr"))
			fmt.Printf("Read MySQL username from opts: %s\n", opts.MySQLOptions.Username)

			jsonData, _ := json.MarshalIndent(opts, "", "  ")
			fmt.Println(string(jsonData))

			return nil
		},
		// 设置命令运行时的参数检查，不需要指定命令行参数。例如：./opsx-center param1 param2
		Args: cobra.NoArgs,
	}

	// 初始化配置函数，在每个命令运行时调用
	cobra.OnInitialize(onInitialize)

	// cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	// 推荐使用配置文件来配置应用，便于管理配置项
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the opsx-center configuration file.")

	return cmd
}
