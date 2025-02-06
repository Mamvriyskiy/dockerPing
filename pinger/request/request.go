package request

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/Mamvriyskiy/dockerPing/pinger/models"
)

const (
	pingerToken = "hsHcmJkmHaJIUzUxMiIsInR5cC3jhmdHJ7H.eyJzdWIiOiIxMjM0NSIsIm5hbWUiOiJKb2huIEdvbGQiLCJhZG1pbiI6dHJ1ZX0K.LIHjWCBORSWMEibq-tnT8ue_deUqZx1K0XxCOXZRrBI"
)

func RequestContainers() ([]models.Container, error) {
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
		return nil, err
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

	var containers []models.Container
    err = json.Unmarshal(body, &containers)
    if err != nil {
        return nil, err
    }

	// fmt.Println(resp.Body, resp.StatusCode)

	return containers, nil

}
