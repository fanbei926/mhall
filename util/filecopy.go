package util

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

func FileCopy(dest string, f string) error {
	fr, err := os.Open(f)
	if err != nil {
		return err
	}
	defer fr.Close()
	reader := bufio.NewReader(fr)

	destFile := filepath.Join(dest, f)
	fw, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer fw.Close()
	writer := bufio.NewWriter(fw)

	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}

	return nil
}
