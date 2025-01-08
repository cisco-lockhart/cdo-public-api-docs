package services

import (
	"os/exec"
)

func InstallNodePackage(packagename string) error {
	_, err := exec.Command("npm", "install", "-g", packagename).Output()
	return err
}

func InstallHomebrewPackage(packagename string) error {
	_, err := exec.Command("brew", "install", packagename).Output()
	return err
}

func InstallPipxPackage(packagename string) error {
	_, err := exec.Command("pipx", "install", packagename, "--include-deps").Output()
	return err
}
