package utilities

import (
	"fmt"
	"os"
	"os/exec"
)

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
