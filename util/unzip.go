package util

import (
	"archive/zip"
	"io"
	"os"
)

// unzip the file
func Unzip(zipFile string) *MhallError {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		mye := New("unzip.go", 13, err.Error())
		return mye
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		if  f.FileInfo().IsDir() {
			if err := os.MkdirAll(f.Name, os.ModePerm); err != nil{
				mye := New("unzip.go", 21, err.Error())
				return mye
			}
			continue
		}

		fr, err := f.Open()
		if err != nil {
			mye := New("unzip.go", 29, err.Error())
			return mye
		}

		fw, err := os.OpenFile(f.Name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			mye := New("unzip.go", 35, err.Error())
			return mye
		}
		_, err = io.Copy(fw, fr)
		if err != nil {
			mye := New("unzip.go", 40, err.Error())
			return mye
		}

		fr.Close()
		fw.Close()
	}
	return nil
}
