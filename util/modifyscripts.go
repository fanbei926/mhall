package util

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"
)

func ModifyScripts(module string, soureFile string) error {
	// 读取 start.sh 文件
	fr, err := os.Open("start.sh")
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("ModifyScripts: Can not find start.sh.")
		} else {
			return errors.New("ModifyScripts: " + err.Error())
		}
	}
	defer fr.Close()

	// 创建新的 start.sh.new 文件
	fw, err := os.OpenFile("start.sh.new", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return errors.New("ModifyScripts: " + err.Error())
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
			return errors.New("ModifyScripts: " + err.Error())
		}

		// 替换脚本中的字段
		line = r.ReplaceAllString(line, soureFile)
		if strings.Contains(line, "tail -F") {
			line = "# " + line
		}
		if _, err = wi.WriteString(line); err != nil {
			return errors.New("ModifyScripts: " + err.Error())
		}

		// 将缓冲区的数据刷进文件中
		if err = wi.Flush(); err != nil {
			return errors.New("ModifyScripts: " + err.Error())
		}

		if err1 == io.EOF {
			break
		}
	}



	return nil
}
