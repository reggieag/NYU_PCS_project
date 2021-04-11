package utilities

import (
	"fmt"
	"os"
	"os/exec"
)

const dockerCompose = "docker-compose"

func StartAPI(file string) error {
	_, err := callCompose(file, "up")
	return err
}

func StopAPI(file string) error {
	cmd, err := callCompose(file, "down")
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func callCompose(file string, command string) (*exec.Cmd, error) {
	cmd := exec.Command(dockerCompose, "-f", file, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}

func RunImage(image string, env map[string]string) error {
	envArguments := make([]string, 0, len(env)*2+5)
	envArguments = append(envArguments, "run", "--rm", "--network", "host")
	for key, value := range env {
		envArguments = append(envArguments, "-e")
		envVariable := fmt.Sprintf("%s=%s", key, value)
		envArguments = append(envArguments, envVariable)
	}
	envArguments = append(envArguments, image)
	cmd := exec.Command("docker", envArguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
