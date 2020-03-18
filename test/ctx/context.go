package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	for {
		time.Sleep(100 * time.Millisecond)
		ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
		ctx = ctx

		runtime.GC()
	}
}
