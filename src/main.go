package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// os.Chdir("../shell")
	params := make([]string, 1)
	params[0] = "php.sh"
	_ = execCommand("/bin/bash", params)
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, "-c", "php.sh")
	//显示运行的命令
	// fmt.Println(cmd.Args)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	return true
}
