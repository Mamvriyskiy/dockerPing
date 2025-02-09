package ping

import (
	"fmt"
	"time"
	"os/exec"
	"strings"
	"github.com/Mamvriyskiy/dockerPing/pinger/models"
)

const (
	countWorkers = 5
)


func workerPingContainer(input, output chan models.Container, i int) {
	for container := range input {
		container.TimePing = time.Now()

		out, err := exec.Command("ping", "-c", "5", "-i", "3", container.ContainerIP).Output()
		if err != nil {
			container.Status = fmt.Sprintf("Ping failed %d", i)
		} else if strings.Contains(string(out), "Destination Host Unreachable") {
			container.Status = "Container is unreachable"
		} else {
			container.Status = "Container is reachable"
		}

		output <- container
	}
}


func CreateWorkersPingContainer(ipContainers []models.Container) []models.Container {
	input := make(chan models.Container)
	output := make(chan models.Container)

	for i := 0; i < countWorkers; i++ {
		go workerPingContainer(input, output, i)
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
