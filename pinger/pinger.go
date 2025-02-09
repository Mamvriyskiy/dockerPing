package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/pinger/request"
	"github.com/Mamvriyskiy/dockerPing/pinger/ping"
)

func main() {
	fmt.Println("Start pinger ...")

	httpServer := RunPinger("8081")

	for {
		logger.Log("Info", fmt.Sprintf("Requesting containers"), nil)

		ipContainers, err := request.RequestContainers()
		if err != nil {
			logger.Log("Error", "Error occurred while requesting containers", err)
			continue
		}

		logger.Log("Info", "Creating workers for pinging containers", nil)
		ping.CreateWorkersPingContainer(ipContainers)

		logger.Log("Info", "Sending status of containers", nil)
		request.SendStatusContainers(ipContainers)
		
		logger.Log("Info", "Sleeping for 20 seconds", nil)
		time.Sleep(20 * time.Second)
	}

	logger.Log("Info", "HTTP server is starting", nil)
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
