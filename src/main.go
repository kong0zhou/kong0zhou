package main

import (
	"fmt"
	"net/http"

	"./common"
	"./controllers"
	"github.com/astaxie/beego/logs"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			return
		}
	}()
	// =========初始化日志文件========
	err := common.InitLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	// ===========初始化配置文件===========
	err = common.InitConf()
	if err != nil {
		logs.Error(err)
		return
	}
	// ===============================
	mux := http.NewServeMux()
	mux.HandleFunc("/show", controllers.Show)
	mux.HandleFunc(`/allFile`, controllers.ErrorHandler(controllers.AllFile))
	logs.Info("http服务器启动，端口：8083")
	err = http.ListenAndServe(":8083", mux)
	if err != nil {
		logs.Error("启动失败", err)
	}
}
