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
				fmt.Println("Can not find the source file.")
				os.Exit(1)
			}
		}

		cfg, err := ini.Load(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		dest := cfg.Section("java").Key("ghall").String()

		switch module {
		case "ghall":
			err := util.Backup(dest, "*.jar")
			if err != nil {
				fmt.Println(err)
			}

			err = util.Deploymodule(sourceFile, dest)
			if err != nil {
				fmt.Println(err)
			}
		case "basic":
			dest := cfg.Section("java").Key("basic").String()
			err := util.Backup(dest, "*.jar")
			if err != nil {
				fmt.Println(err)
			}

			err = util.Deploymodule(sourceFile, dest)
			if err != nil {
				fmt.Println(err)
			}
		case "hall":
			dest := cfg.Section("java").Key("hall").String()
			err := util.Backup(dest, "*.jar")
			if err != nil {
				fmt.Println(err)
			}

			err = util.Deploymodule(sourceFile, dest)
			if err != nil {
				fmt.Println(err)
			}
		case "dscrm":
			dest := cfg.Section("java").Key("dscrm").String()
			err := util.Backup(dest, "*.jar")
			if err != nil {
				fmt.Println(err)
			}

			err = util.Deploymodule(sourceFile, dest)
			if err != nil {
				fmt.Println(err)
			}
		case "gh":
			dest := cfg.Section("nginx").Key("gh").String()
			err := util.Backup(dest, "gh")
			if err != nil {
				fmt.Println(err)
			}

			err = util.Deploymodule(sourceFile, path.Join(dest, module))
			if err != nil {
				fmt.Println(err)
			}
		case "Ghall":
			dest := cfg.Section("nginx").Key("Ghall").String()
			err := util.Backup(dest, "Ghall")
			if err != nil {
				fmt.Println(err)
			}

			err = util.Deploymodule(sourceFile, path.Join(dest, module))
			if err != nil {
				fmt.Println(err)
			}
		case "Video":
			dest := cfg.Section("nginx").Key("Video").String()
			err := util.Backup(dest, "Video")
			if err != nil {
				fmt.Println(err)
			}

			err = util.Deploymodule(sourceFile, dest)
			if err != nil {
				fmt.Println(err)
			}
		case "Video2":
			dest := cfg.Section("nginx").Key("Video2").String()
			err := util.Backup(dest, "Video2")
			if err != nil {
				fmt.Println(err)
			}

			err = util.Deploymodule(sourceFile, dest)
			if err != nil {
				fmt.Println(err)
			}
		}

		// 切换到目标目录，如果有 start.sh 脚本，进行修改并执行
		err = os.Chdir(dest)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _, err = os.Stat("start.sh"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = util.ModifyScripts(module, sourceFile)
		if err != nil {
			fmt.Println(err)
		} else {
			err = util.ExecuteShell()
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("Deploy " + module + " done.")
	},
}
