package services

import (
	"github.com/cisco-lockhart/fcm-api-docs-generator/models"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadConfig(filename string) *models.Config {
	// Load the configuration file
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	config := models.Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
