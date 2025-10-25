package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewOpsXCenterCommand 创建一个 *cobra.Command 对象，用于启动应用程序.
func NewOpsXCenterCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "opsx-center",
		// 命令的简短描述
		Short: "devops center",
		Long: `devops center, designed to help beginners quickly
		learn Go project development.`,
		// 命令出错时，不打印帮助信息。设置为 true 可以确保命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello opsx!")
			return nil
		},
		// 设置命令运行时的参数检查，不需要指定命令行参数。例如：./opsx-center param1 param2
		Args: cobra.NoArgs,
	}
	return cmd
}
