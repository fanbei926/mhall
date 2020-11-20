package util

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func Deploymodule(source string, dest string) error {
	// 拷贝文件到目标目录
	currDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// 如果不存在dest 则创建
	if _, err := os.Stat(dest); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(dest, os.ModePerm); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// 检查源文件是文件夹还是文件
	s, err := os.Stat(path.Join(currDir, source))
	if err != nil {
		return err
	}
	// 如果是文件夹，不需要解压直接rename过去
	if s.IsDir() {
		err = os.Rename(source, path.Join(dest, source))
		if err != nil {
			return err
		}
	} else {
		// todo: 其实所有文件都需要rename过去，而不是copy过去
		err = FileCopy(dest, source)
		if err != nil {
			return err
		}
	}

	// 切换到目标目录，进行解压
	os.Chdir(dest)
	suffix := strings.Split(source, ".")
	if len(suffix) == 2 && suffix[1] == "zip" {
		err := Unzip(source)
		if err != nil {
			return err
		}
	} else {
		fmt.Print("Do not need unzip.")
	}

	if err := os.Chdir(currDir); err != nil {
		return err
	}
	return nil
}
