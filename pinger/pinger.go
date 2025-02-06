package main

import (
	"fmt"
	// "github.com/Mamvriyskiy/dockerPing/logger"
)

const (
	pingerToken = "hsHcmJkmHaJIUzUxMiIsInR5cC3jhmdHJ7H.eyJzdWIiOiIxMjM0NSIsIm5hbWUiOiJKb2huIEdvbGQiLCJhZG1pbiI6dHJ1ZX0K.LIHjWCBORSWMEibq-tnT8ue_deUqZx1K0XxCOXZRrBI"
)

func main() {
	fmt.Println("Start pinger ...")

	httpServer := RunPinger("8081")

	for {
		
	}

	httpServer.ListenAndServe()
}

func RunPinger(port string) *http.Server {
	return &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
}
