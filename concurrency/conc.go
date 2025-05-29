package concurrency

import (
	"fmt"
	"os"
	"path/filepath"
)

// creates random files with random texts s is number of strings to add ,n is no of files
func CreateFiles(n, nt int, s string) error {

	if n == 0 {
		return fmt.Errorf("files can't be zero ")
	}
	err := os.Mkdir(s, 0777)
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		basePath, err := os.Getwd()
		if err != nil {
			return err
		}
		path := filepath.Join(basePath, s, RandString("file_"))
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		for j := 0; j < nt; j++ {
			if _, err := f.Write([]byte(s)); err != nil {
				return err
			}
		}
	}
	return nil

}
