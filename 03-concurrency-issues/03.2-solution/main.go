package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var countVisitors int64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&countVisitors, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Hello, you are the %d visitor", countVisitors)))
	})

	http.ListenAndServe(":8080", nil)
}
