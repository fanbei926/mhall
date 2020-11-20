package util

import (
	"archive/zip"
	"io"
	"os"
)

// unzip the file
func Unzip(zipFile string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		if  f.FileInfo().IsDir() {
			if err := os.MkdirAll(f.Name, os.ModePerm); err != nil{
				return err
			}
			continue
		}

		fr, err := f.Open()
		if err != nil {
			return err
		}

		fw, err := os.OpenFile(f.Name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return  err
		}
		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}

		fr.Close()
		fw.Close()
	}
	return nil
}
