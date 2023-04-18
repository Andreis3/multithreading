package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var countVisitors int64 = 0

func main() {
	mutex := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		countVisitors++
		mutex.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Hello, you are the %d visitor", countVisitors)))
	})

	http.ListenAndServe(":8080", nil)
}
