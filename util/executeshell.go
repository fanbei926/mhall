package util

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteShell() error {
	if err := os.Remove("start.sh"); err != nil {
		fmt.Println(err)
	}
	os.Rename("./start.sh.new", "./start.sh")
	comm := exec.Command("bash", "-c", "start.sh")
	stdErr, _ := comm.StderrPipe()
	if err := comm.Start(); err != nil {
		return err
	}

	for  {
		buf := make([]byte, 1024)
		_, err := stdErr.Read(buf)
		if err != nil {
			return err
		}
		fmt.Println(buf)
	}

	if err := comm.Wait(); err != nil {
		return  err
	}

	return nil
}
