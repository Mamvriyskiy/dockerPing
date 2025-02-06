package request

import (
	"time"
	"net/http"
	"io/ioutil"
)

func requestContainers() {
	requestURL := "http://localhost:8000/api/ping"
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		// fmt.Printf("error making http request: %s\n", err)
		// os.Exit(1)
	}

	req.Header.Set("Authorization", "Bearer " + pingerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// fmt.Println(resp.Body, resp.StatusCode)
	time.Sleep(10 * time.Second)

	body, _ := ioutil.ReadAll(resp.Body)

	pingContainers(body)
	// fmt.Println("Response:", string(body))
}
