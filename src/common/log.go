package common

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego/logs"
)

func InitLogger() error {

	logs.SetLogFuncCallDepth(3)    //调用层级
	logs.EnableFuncCallDepth(true) //输出文件名和行号
	if ConfViper.GetBool(`production`) {
		logs.Async() //提升性能, 可以设置异步输出
	}
	config := make(map[string]interface{})
	config["filename"] = `./logs/log.log`
	config["level"] = logs.LevelDebug
	config["perm"] = "0777"

	configStr, err := json.Marshal(config)
	if err != nil {
		log.Fatal("initLogger failed, marshal err:", err)
		return err
	}
	if !ConfViper.GetBool(`production`) {
		err = logs.SetLogger(logs.AdapterConsole, "") //控制台输出
		if err != nil {
			log.Fatal("SetLogger failed, err:", err)
			return err
		}
	}
	err = logs.SetLogger(logs.AdapterFile, string(configStr)) //文件输出

	if err != nil {
		log.Fatal("SetLogger failed, err:", err)
		return err
	}
	return nil
}
