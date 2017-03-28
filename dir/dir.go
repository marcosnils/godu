package dir

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Walk(dir string, size chan int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			Walk(subdir, size)
		} else {
			size <- entry.Size()
		}
	}

}

func dirents(dir string) []os.FileInfo {
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	return dirFiles

}
