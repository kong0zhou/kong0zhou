package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// os.Chdir("../shell")
	params := make([]string, 1)
	params[0] = "php.sh"
	_ = execCommand("bash", params)
	fmt.Println("21564564")
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe() //接收命令在控制行里输出的数据（字符串）
	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容(即打印到控制的数据)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println("hey:", line)
	}

	cmd.Wait()
	return true
}
