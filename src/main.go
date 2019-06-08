package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// os.Chdir("../shell")
	params := make([]string, 2)
	params[0] = "-c"
	params[1] = `
	echo "shdfkjshdf"
	PHP=$(pwd)
	docker run --name php \
		--rm \
		-it \
		-v $PHP/code:/code \
		php \
		php /code/index.php`
	// params[1] = "php.sh"
	b := execCommand("bash", params)
	if b {
		fmt.Println("454545")
	}
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	fmt.Println(string(out))
	return true
}
