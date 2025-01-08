package services

import (
	"encoding/json"
	"errors"
	"github.com/Masterminds/semver/v3"
	"github.com/cisco-lockhart/fcm-api-docs-generator/models"
	"io"
	"net/http"
)

const pypiPackageName = "cdo_sdk_python"

func GetCurrentVersion(openApiVersionStr string) (*string, error) {
	resp, err := http.Get("https://pypi.org/pypi/" + pypiPackageName + "/json")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch URL: " + resp.Request.URL.String() + ", status code: " + resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	pypiVersionInfo := models.PypiVersionInfo{}
	if err = json.Unmarshal(data, &pypiVersionInfo); err != nil {
		return nil, err
	}

	if pypiVersionInfo.Info.Version == "" {
		return &openApiVersionStr, nil
	}

	openApiVersion, err := semver.NewVersion(openApiVersionStr)
	if err != nil {
		return nil, err
	}
	pypiVersion, err := semver.NewVersion(pypiVersionInfo.Info.Version)
	if err != nil {
		return nil, err
	}
	if openApiVersion.Major() == pypiVersion.Major() && openApiVersion.Minor() == pypiVersion.Minor() {
		newVersion := pypiVersion.IncPatch()
		newVersionStr := newVersion.String()
		return &newVersionStr, nil
	}

	return &openApiVersionStr, nil
}
