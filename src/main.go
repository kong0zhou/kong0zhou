package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// os.Chdir("../shell")
	params := make([]string, 2)
	params[0] = "-c"
	params[1] = "php.sh"
	_ = execCommand("/bin/bash", params)
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	// fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe() //接收命令在控制行里输出的数据（字符串）

	if err != nil {
		fmt.Println(err)
		return false
	}
	cmd.Start()

	cmd.Wait()
	reader, _ := ioutil.ReadAll(stdout)
	s := string(reader) + "454564"
	fmt.Println(s)

	return true
}
