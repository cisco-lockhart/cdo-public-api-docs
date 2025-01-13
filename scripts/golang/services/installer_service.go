package services

func InstallNodePackage(packagename string) error {
	return InstallNodePackageGloballyOrLocally(packagename, true)
}

func InstallNodePackageGloballyOrLocally(packagename string, global bool) error {
	var err error
	if global {
		_, err = ExecCommand("npm", "install", "-g", packagename).Output()
	} else {
		_, err = ExecCommand("npm", "install", packagename).Output()
	}
	return err
}

func InstallHomebrewPackage(packagename string) error {
	_, err := ExecCommand("brew", "install", packagename).Output()
	return err
}

func InstallPipxPackage(packagename string) error {
	_, err := ExecCommand("pipx", "install", packagename, "--include-deps").Output()
	return err
}
