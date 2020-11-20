package util

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

// dest: 程序所在目录
// suffix: 需要备份的文件所包涵的后缀
func Backup(dest string, suffix string) error {
	currDate := time.Now().Format("20060102_150405")
	currDir, err := os.Getwd()
	if err != nil {
		return nil
	}

	// 切换到backup目录
	if err := os.Chdir(dest); err != nil {
		return errors.New(dest + " folder not found. Later it will be created.")
	}
	// 如果不存在backup 则创建
	if _, err := os.Stat("Backup"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Backup does not exist. Now create it.")
			if err = os.Mkdir("Backup", os.ModePerm); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// 找到所有包含suffix的文件，将其移动到backup文件夹中
	oldFiles, err := filepath.Glob(suffix)
	if err != nil {
		return err
	}

	for _, f := range oldFiles {
		oldpath := path.Join(".", f)
		newpath := path.Join("Backup", f + "." + currDate)
		if err := os.Rename(oldpath, newpath); err != nil {
			return err
		}
	}

	if err := os.Chdir(currDir); err != nil {
		return err
	}
	return nil
}