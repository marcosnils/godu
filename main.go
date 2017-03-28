package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type ReqParam struct {
	Dir string
}

func main() {

	http.HandleFunc("/size", func(rw http.ResponseWriter, req *http.Request) {
		size := make(chan int64)
		var p ReqParam
		json.NewDecoder(req.Body).Decode(&p)
		fmt.Println(p.Dir)

		go func() {
			walkDir(p.Dir, size)
			close(size)
		}()
		var total int64
		for s := range size {
			total += s
		}

		fmt.Printf("%.2f GB\n", float32(total)/1e9)
	})

	fmt.Println("Inicializando servidor")
	fmt.Println(http.ListenAndServe(":8080", nil))

}

func walkDir(dir string, size chan int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, size)
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
