package util

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

// comment: 如果是0，表示注释；如果是1，表示取消注释
func ModifyScripts(module string, soureFile string, comment int8) *MhallError {
	// 读取 start.sh 文件
	fr, err := os.Open("start.sh")
	if err != nil {
		if os.IsNotExist(err) {
			mye := New("modifyscripts.go", 18, "Can not find start.sh.")
			return mye
		} else {
			mye := New("modifyscripts.go", 21, err.Error())
			return mye
		}
	}
	defer fr.Close()

	// 创建新的 start.sh.new 文件
	fw, err := os.OpenFile("start.sh.new", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		mye := New("modifyscripts.go", 30, err.Error())
		return mye
	}
	defer fw.Close()

	rd := bufio.NewReader(fr)
	wi := bufio.NewWriter(fw)

	temp := module + `.*jar`
	r,_ := regexp.Compile(temp)
	rd.ReadLine()
	for {
		line, err1 := rd.ReadString('\n')
		if err1 != nil && err1 != io.EOF{
			mye := New("modifyscripts.go", 44, err.Error())
			return mye
		}

		// 一般如果需要注释，肯定同时也需要替换jar包名称
		if comment == 0 {
			line = r.ReplaceAllString(line, soureFile)
			if strings.Contains(line, "tail -F") {
				// 替换脚本中的字段
				line = "# " + line
			}
		} else {
			line = strings.Replace(line, "# ", "", -1)
		}

		if _, err = wi.WriteString(line); err != nil {
			mye := New("modifyscripts.go", 60, err.Error())
			return mye
		}

		// 将缓冲区的数据刷进文件中
		if err = wi.Flush(); err != nil {
			mye := New("modifyscripts.go", 66, err.Error())
			return mye
		}

		if err1 == io.EOF {
			break
		}
	}

	// 替换掉原有的 start.sh
	if err := os.Remove("start.sh"); err != nil {
		mye := New("modifyscripts.go", 30, err.Error())
		return mye
	}
	os.Rename("./start.sh.new", "./start.sh")

	return nil
}
