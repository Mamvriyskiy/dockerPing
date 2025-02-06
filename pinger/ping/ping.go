package ping

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/Mamvriyskiy/dockerPing/pinger/models"
)

func PingContainers(ipContainers []models.Container) {
	fmt.Println(ipContainers)
	for _, ip := range ipContainers {
		out, _ := exec.Command("ping", ip.ContainerIP, "-c 5", "-i 3").Output()
		if strings.Contains(string(out), "Destination Host Unreachable") {
			fmt.Println("TANGO DOWN")
		} else {
			fmt.Println("IT'S ALIVEEE")
		}
	}

	// if strings.Contains(outputStr, "0% packet loss") {
	// 	return true, nil
	// }

	// return false, nil
}
