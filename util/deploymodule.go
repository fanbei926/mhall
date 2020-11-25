package util

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func Deploymodule(source string, dest string) *MhallError {
	// 拷贝文件到目标目录
	currDir, err := os.Getwd()
	if err != nil {
		mye := New("deploymodule.go", 14, err.Error())
		return mye
	}
	// 如果不存在dest 则创建
	if _, err := os.Stat(dest); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(dest, os.ModePerm); err != nil {
				mye := New("deploymodule.go", 21, err.Error())
				return mye
			}
		} else {
			mye := New("deploymodule.go", 25, err.Error())
			return mye
		}
	}

	// 检查源文件是文件夹还是文件
	s, err := os.Stat(path.Join(currDir, source))
	if err != nil {
		mye := New("deploymodule.go", 33, err.Error())
		return mye
	}
	// 如果是文件夹，不需要解压直接rename过去
	if s.IsDir() {
		err = os.Rename(source, path.Join(dest, source))
		if err != nil {
			mye := New("deploymodule.go", 40, err.Error())
			return mye
		}
	} else {
		// todo: 其实所有文件都需要rename过去，而不是copy过去
		err = FileCopy(dest, source)
		if err != nil {
			mye := New("deploymodule.go", 47, err.Error())
			return mye
		}
	}

	// 切换到目标目录，进行解压
	os.Chdir(dest)
	suffix := strings.Split(source, ".")
	if len(suffix) == 2 && suffix[1] == "zip" {
		err := Unzip(source)
		if err != nil {
			mye := New("deploymodule.go", 58, err.Error())
			return mye
		}
	} else {
		fmt.Println("Do not need unzip.")
	}

	if err := os.Chdir(currDir); err != nil {
		mye := New("deploymodule.go", 66, err.Error())
		return mye
	}
	return nil
}
