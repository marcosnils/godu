package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marcosnils/godu1/dir"
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
			dir.Walk(p.Dir, size)
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
