package request

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/pinger/models"
)

const (
	pingerToken = "hsHcmJkmHaJIUzUxMiIsInR5cC3jhmdHJ7H.eyJzdWIiOiIxMjM0NSIsIm5hbWUiOiJKb2huIEdvbGQiLCJhZG1pbiI6dHJ1ZX0K.LIHjWCBORSWMEibq-tnT8ue_deUqZx1K0XxCOXZRrBI"
)

func RequestContainers() ([]models.Container, error) {
	requestURL := "http://backend:8000/api/ping"
	// requestURL := "http://localhost:8000/api/ping"
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		logger.Log("Error", "Error creating HTTP request", err, fmt.Sprintf("requestURL = %s", requestURL))
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + pingerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Log("Error", "Error making HTTP request", err, fmt.Sprintf("requestURL = %s", requestURL))
		return nil, err
	}
	defer resp.Body.Close()

	logger.Log("Info", fmt.Sprintf("Received response with status: %d", resp.StatusCode), nil)
	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("unexpected response status: %d", resp.StatusCode)
		logger.Log("Error", "Unexpected response status", nil)
    	err = fmt.Errorf(errMsg)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log("Error", "Error reading response body", err, nil)
		return nil, err
	}

	var containers []models.Container
	err = json.Unmarshal(body, &containers)
	if err != nil {
		logger.Log("Error", "Error unmarshaling response body", err, nil)
		return nil, err
	}

	logger.Log("Info", "Successfully received and parsed containers", nil)
	return containers, nil
}

func SendStatusContainers(ipContainers []models.Container) error {
	logger.Log("Info", "Starting to send container status", nil)

	requestURL := "http://backend:8000/api/pinger"

	jsonData, err := json.Marshal(ipContainers)
	if err != nil {
		logger.Log("Error", "Error marshaling JSON", err, nil)
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	logger.Log("Info", "Sending the following JSON data", nil)

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Log("Error", "Error creating HTTP request", err, requestURL)
		return fmt.Errorf("error creating HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+pingerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Log("Error", "Error making HTTP request", err, requestURL)
		return fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	logger.Log("Info", fmt.Sprintf("Received response with status: %d", resp.StatusCode), nil)

	if resp.StatusCode != http.StatusOK {
		logger.Log("Error", fmt.Sprintf("unexpected response status: %d", resp.StatusCode), nil)
		return fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	logger.Log("Info", "Successfully sent container status", nil)

	return nil
}
