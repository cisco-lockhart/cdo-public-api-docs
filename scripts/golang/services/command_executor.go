package services

import "os/exec"

type CommandExecutor interface {
	Output() ([]byte, error)
}

var ExecCommand = func(name string, args ...string) CommandExecutor {
	return exec.Command(name, args...)
}
