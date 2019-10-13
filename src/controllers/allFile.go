package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"../common"
	"github.com/astaxie/beego/logs"
)

func AllFile(w http.ResponseWriter, r *http.Request) {
	reply, err := NewReplyProto(`GET`, `/allFile`)
	if err != nil {
		logs.Error(err)
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
	dirPath := common.ConfViper.GetString(`dirPath`)
	if dirPath == `` {
		err = fmt.Errorf(`dirPath of config file is not found`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	dirExists, err := common.PathExists(dirPath)
	if err != nil {
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	if !dirExists {
		err = fmt.Errorf(`dirPath is not exists`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}

	allFile := make([]string, 0)
	err = filepath.Walk(dirPath,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				logs.Error(err)
				return err
			}
			if f == nil {
				logs.Error(err)
				return err
			}
			if path == "" {
				err = fmt.Errorf("path is null")
				logs.Error(err)
				return err
			}
			//判断是否是文件夹，如果是文件夹，直接返回，不读取
			if f.IsDir() {
				return nil
			}
			allFile = append(allFile, path)
			return nil
		})
	if err != nil {
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	err = reply.SuccessResp(allFile, w)
	if err != nil {
		logs.Error(err)
		return
	}
}
