package services

import (
	"fmt"
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/models"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
)

func LoadConfig(url string) (*models.Config, error) {
	// Load the configuration file
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("failed to fetch URL: %s, status code: %d", url, resp.StatusCode))
	}

	// Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	config := models.Config{}
	err = yaml.Unmarshal(data, &config)
	fmt.Printf("%v\n", config.Info)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
