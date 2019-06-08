package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// os.Chdir("../shell")
	params := make([]string, 1)
	params[0] = "php.sh"
	// params[1] = "php.sh"
	b := execCommand("bash", params)
	if b {
		fmt.Println("454545")
	}
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
