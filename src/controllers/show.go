package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"../common"

	"github.com/astaxie/beego/logs"
)

func Show(w http.ResponseWriter, r *http.Request) {
	sse, err := NewSse(w)
	if err != nil {
		logs.Error(err)
		return
	}
	reply, err := NewReplyProto(`GET`, `/show`)
	if err != nil {
		logs.Error(err)
		return
	}
	if r.Method != "GET" {
		logs.Info(`somebody do not use GET method`)
		logs.Info(r.Method)
		err := reply.SseError(`you must use GET method`, sse)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	logs.Info(r.Method)
	data := r.URL.Query().Get("q")
	logs.Info(data)
	// var filePath string
	var req ReqProto
	err = json.Unmarshal([]byte(data), &req)
	if err != nil {
		logs.Error(err)
		err = reply.SseError(err.Error(), sse)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	filePath, ok := req.Data.(string)
	if !ok {
		err = fmt.Errorf(`req.Data must be string`)
		logs.Error(err)
		err = reply.SseError(err.Error(), sse)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	if filePath == `` {
		err = fmt.Errorf(`filePath is null`)
		logs.Error(err)
		err = reply.SseError(err.Error(), sse)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	fileExists, err := common.PathExists(filePath)
	if err != nil {
		logs.Error(err)
		err = reply.SseError(err.Error(), sse)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	if !fileExists {
		err = fmt.Errorf(`file is not exist`)
		logs.Error(err)
		err = reply.SseError(err.Error(), sse)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	logs.Info(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		logs.Error(err)
		err = reply.SseError(err.Error(), sse)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		logs.Error(err)
		err = reply.SseError(err.Error(), sse)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	// 指定返回前端的字节数
	var lmz int64
	if fi.Size() < logMaxSize {
		lmz = fi.Size()
	} else {
		lmz = logMaxSize
	}
	_, err = file.Seek(-lmz, 2)
	if err != nil {
		logs.Error(err)
		err = reply.SseError(err.Error(), sse)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadBytes('\n')
		select {
		case <-sse.IsClosed:
			return
		default:
			if err == io.EOF {
				if len(line) != 0 {
					err = reply.SseSuccess(string(line), sse)
					if err != nil {
						logs.Error(err)
						return
					}
				}
				time.Sleep(500 * time.Millisecond)
				continue
			} else if err != nil {
				logs.Error(err)
				err = reply.SseError(err.Error(), sse)
				if err != nil {
					logs.Error(err)
					return
				}
				return
			}
			// fmt.Println(string(line))
			err = reply.SseSuccess(string(line), sse)
			if err != nil {
				logs.Error(err)
				return
			}
		}
	}
}

// func FileMonitoring(file *os.File, ctx context.Context, w http.ResponseWriter) (errChan chan error) {
// 	errChan = make(chan error, 1)
// 	if file == nil {
// 		err := fmt.Errorf("file is nil")
// 		logs.Error(err)
// 		errChan <- err
// 		return
// 	}
// 	if
// 	rd := bufio.NewReader(file)
// 	file.Seek(0, 2)
// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			default:
// 				line, err := rd.ReadBytes('\n')
// 				if err == io.EOF {
// 					time.Sleep(500 * time.Millisecond)
// 					continue
// 				} else if err != nil {
// 					logs.Error(err)
// 					errChan <- err
// 					return
// 				}
// 				lineChan <- string(line)
// 			}
// 		}
// 	}()
// 	return
// }
