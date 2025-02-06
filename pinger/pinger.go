package main

import (
	"fmt"
	"time"
	"net/http"
	// "github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/pinger/request"
	"github.com/Mamvriyskiy/dockerPing/pinger/ping"
)

func main() {
	fmt.Println("Start pinger ...")

	httpServer := RunPinger("8081")

	//добавить время засыпания из конфига, количество воркеров
	for {
		ipContainers, err := request.RequestContainers()
		if err != nil {
			// error log
			continue
		}

		ping.PingContainers(ipContainers)

		time.Sleep(20 * time.Second)
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
