package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

// dest: 程序所在目录
// suffix: 需要备份的文件所包涵的后缀
func Backup(dest string, suffix string) *MhallError {
	currDate := time.Now().Format("20060102_150405")
	currDir, err := os.Getwd()
	if err != nil {
		mye := New("backup.go", 17, err.Error())
		return mye
	}

	// 切换到backup目录
	if err := os.Chdir(dest); err != nil {
		mye := New("backup.go", 23, dest + " folder not found. Later it will be created.")
		return mye
	}
	// 如果不存在backup 则创建
	if _, err := os.Stat("Backup"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Backup does not exist. Now create it.")
			if err = os.Mkdir("Backup", os.ModePerm); err != nil {
				mye := New("backup.go", 31, err.Error())
				return mye
			}
		} else {
			mye := New("backup.go", 37, err.Error())
			return mye
		}
	}

	// 找到所有包含suffix的文件，将其移动到backup文件夹中
	oldFiles, err := filepath.Glob(suffix)
	if err != nil {
		mye := New("backup.go", 43, err.Error())
		return mye
	}

	for _, f := range oldFiles {
		oldpath := path.Join(".", f)
		newpath := path.Join("Backup", f + "." + currDate)
		if err := os.Rename(oldpath, newpath); err != nil {
			mye := New("backup.go", 53, err.Error())
			return mye
		}
	}

	if err := os.Chdir(currDir); err != nil {
		mye := New("backup.go", 56, err.Error())
		return mye
	}
	return nil
}