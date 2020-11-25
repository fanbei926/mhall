package util

import (
	"fmt"
	"os/exec"
)

func ExecuteShell(scriptfile string) error {

	comm := exec.Command("bash", "-c", scriptfile)
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
