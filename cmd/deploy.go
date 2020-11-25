package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
	"mhall/util"
	"os"
	"path"
)

var conf string
var module string

func init()  {
	cobra.OnInitialize()
	deployCmd.Flags().StringVarP(&module, "module", "m", "", "The module you want to deploy")
	deployCmd.Flags().StringVarP(&conf, "config", "f", ".\\conf\\basic.ini", "The basic config's file")
	deployCmd.MarkFlagRequired("module")
}

var deployCmd = &cobra.Command{
	Use: "deploy",
	Short: "deploy function",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceFile := args[0]
		if _, err := os.Stat(sourceFile); err != nil {
			if os.IsNotExist(err) {
				mye := util.New("deploy.go", 31, err.Error())
				fmt.Println(mye.Error())
				os.Exit(1)
			}
		}

		cfg, err := ini.Load(conf)
		if err != nil {
			mye := util.New("deploy.go", 38, err.Error())
			fmt.Println(mye.Error())
			os.Exit(1)
		}

		// 根据不同的 module 进行相应的操作
		switch module {
		case "ghall", "basic", "hall", "dscrm":
			// 获取配置
			dest := cfg.Section("java").Key(module).String()
			if dest == "" {
				mye := util.New("deploy.go", 49, "Path is blank. Please check.")
				fmt.Println(mye.Error())
				os.Exit(1)
			}

			mye := util.Backup(dest, "*.jar")
			if mye != nil {
				fmt.Println(mye.Error())
				os.Exit(1)
			}

			mye = util.Deploymodule(sourceFile, dest)
			if mye != nil {
				fmt.Println(mye.Error())
				os.Exit(1)
			}

			// 切换到目标目录，如果有 start.sh 脚本，进行修改并执行
			err = os.Chdir(dest)
			if err != nil {
				mye := util.New("deploy.go", 69, err.Error())
				fmt.Println(mye.Error())
				os.Exit(1)
			}
			if _, err = os.Stat("start.sh"); err != nil {
				mye := util.New("deploy.go", 74, err.Error())
				fmt.Println(mye.Error())
				os.Exit(1)
			}

			// 注释、更改jar包名称
			err = util.ModifyScripts(module, sourceFile, 0)
			if err != nil {
				mye := util.New("deploy.go", 82, err.Error())
				fmt.Println(mye.Error())
				os.Exit(1)
			}

			err = util.ExecuteShell("stop.sh")
			if err != nil {
				mye := util.New("deploy.go", 89, err.Error())
				fmt.Println(mye.Error())
				os.Exit(1)
			}
			err = util.ExecuteShell("start.sh")
			if err != nil {
				mye := util.New("deploy.go", 95, err.Error())
				fmt.Println(mye.Error())
				os.Exit(1)
			}
			// 取消注释
			err = util.ModifyScripts(module, sourceFile, 1)
			if err != nil {
				mye := util.New("deploy.go", 101, err.Error())
				fmt.Println(mye.Error())
				os.Exit(1)
			}

		case "gh", "Ghall":
			// 获取配置
			dest := cfg.Section("nginx").Key(module).String()
			if dest == "" {
				mye := util.New("deploy.go", 111, "Path is blank. Please check.")
				fmt.Println(mye.Error())
				os.Exit(1)
			}

			mye := util.Backup(dest, module)
			if mye != nil {
				fmt.Println(mye.Error())
				os.Exit(1)
			}

			mye = util.Deploymodule(sourceFile, path.Join(dest, module))
			if mye != nil {
				fmt.Println(mye.Error())
				os.Exit(1)
			}

		case "Video", "Video2":
			// 获取配置
			dest := cfg.Section("nginx").Key(module).String()
			if dest == "" {
				mye := util.New("deploy.go", 132, "Path is blank. Please check.")
				fmt.Println(mye.Error())
				os.Exit(1)
			}

			mye := util.Backup(dest, module)
			if mye != nil {
				fmt.Println(mye.Error())
				os.Exit(1)
			}

			mye = util.Deploymodule(sourceFile, dest)
			if mye != nil {
				fmt.Println(mye.Error())
				os.Exit(1)
			}
		}

		fmt.Println("Deploy " + module + " done.")
	},
}
