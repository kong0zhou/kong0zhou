package controllers

import (
	"fmt"
	"net/http"

	"../common"
	"github.com/astaxie/beego/logs"
)

type fileData struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func AllFile(w http.ResponseWriter, r *http.Request) {
	reply, err := NewReplyProto(`GET`, `/allFile`)
	if err != nil {
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	if r.Method != "GET" {
		logs.Info(`somebody do not use GET method`)
		err := reply.ErrorResp(`you must use GET method`, w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	fileNames := common.ConfViper.GetStringSlice(`fileName`)
	if fileNames == nil || len(fileNames) == 0 {
		err = fmt.Errorf(`fileNames of config file is not found`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	filePath := common.ConfViper.GetStringSlice(`filePath`)
	if filePath == nil || len(filePath) == 0 {
		err = fmt.Errorf(`filePath of config file is not found`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	fileDatas := make([]fileData, 0)
	for i, v := range fileNames {
		var fileData fileData
		fileData.Name = v
		fileData.Path = filePath[i]
		fileDatas = append(fileDatas, fileData)
	}
	err = reply.SuccessResp(fileDatas, w)
	if err != nil {
		logs.Error(err)
		return
	}
}
