package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "php.sh")
	//显示运行的命令
	// fmt.Println(cmd.Args)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, "-c", "php.sh")
	//显示运行的命令
	// fmt.Println(cmd.Args)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(out))
	return true
}
