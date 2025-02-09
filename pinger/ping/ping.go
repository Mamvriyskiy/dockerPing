package ping

import (
	"time"
	"os/exec"
	"strings"
	"github.com/Mamvriyskiy/dockerPing/pinger/models"
)

const (
	countWorkers = 5
)

func workerPingContainer(input, output chan models.Container) {
	for container := range input {
		container.TimePing = time.Now()

		out, err := exec.Command("ping", "-c", "5", "-i", "3", container.ContainerIP).Output()
		if err != nil {
			container.Status = "Ping failed: " + err.Error()
		} else if strings.Contains(string(out), "Destination Host Unreachable") {
			container.Status = "Container is unreachable"
		} else if strings.Contains(string(out), "100% packet loss") {
			container.Status = "Container is offline (no response)"
		} else if strings.Contains(string(out), "0% packet loss") {
			container.Status = "Container is reachable and online"
		} else {
			container.Status = "Container status unknown: " + string(out)
		}

		output <- container
	}
}


func CreateWorkersPingContainer(ipContainers []models.Container) []models.Container {
	input := make(chan models.Container)
	output := make(chan models.Container)

	for i := 0; i < countWorkers; i++ {
		go workerPingContainer(input, output)
	}

	go func() {
		for _, ip := range ipContainers {
			input <- ip
		}
		close(input)
	}()

	for i := 0; i < len(ipContainers); i++ {
		ipContainers[i] = <-output
	}

	return ipContainers
}
