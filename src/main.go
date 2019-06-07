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

	reader, _ := ioutil.ReadAll(stdout)

	//实时循环读取输出流中的一行内容(即打印到控制的数据)
	fmt.Println(string(reader))

	cmd.Wait()
	return true
}
