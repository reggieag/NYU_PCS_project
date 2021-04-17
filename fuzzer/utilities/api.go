package utilities

import (
	"os"
	"os/exec"
)

func StartAPI(file string) error {
	cmd := exec.Command(file, "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func StopAPI(file string) error {
	cmd := exec.Command(file, "stop")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
