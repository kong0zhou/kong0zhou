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
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	// fmt.Println(cmd.Args)
	out, _ := cmd.Output()
	fmt.Println(string(out))
	return true
}
