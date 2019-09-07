package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"../common"

	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// EnableCompression: false,
}

func Show(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(`panic:`, err)
			return
		}
	}()
	logs.Info("123")
	// 解决 websocket: request origin not allowed by Upgrader.CheckOrigin
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Error(err)
		return
	}
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pathChan, isClose, errChan := wsRead(conn, ctx)
	tailChan, errChan2 := wsDispatch(pathChan, ctx)
	errChan3 := wsWrite(conn, tailChan, ctx)
	select {
	case err := <-errChan:
		logs.Error(err)
		return
	case err := <-errChan2:
		logs.Error(err)
		return
	case err := <-errChan3:
		logs.Error(err)
		return
	case <-isClose:
		logs.Info(`is close`)
		return
	}
}

func wsRead(conn *websocket.Conn, ctx context.Context) (pathChan chan string, isClose chan bool, errChan chan error) {
	isClose = make(chan bool)
	pathChan = make(chan string, 10)
	errChan = make(chan error)
	if conn == nil {
		err := fmt.Errorf(`conn is null`)
		logs.Error(err)
		errChan <- err
		return
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.Error(err)
				return
			}
		}()
		var req ReqProto
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := conn.ReadJSON(&req)
				if err != nil {
					if websocket.IsCloseError(err, 1001) {
						logs.Info(`websocket is close`, err)
						isClose <- true
						return
					}
					logs.Error(err)
					errChan <- err
					return
				}
				path, ok := req.Data.(string)
				logs.Info(path)
				if !ok {
					err := fmt.Errorf(`path must be string`)
					logs.Error(err)
					errChan <- err
					return
				}
				pathChan <- path
			}
		}
	}()
	return
}

func wsDispatch(pathChan chan string, ctx context.Context) (tailChan chan *tail.Tail, errChan chan error) {
	tailChan = make(chan *tail.Tail, 10)
	errChan = make(chan error)
	if pathChan == nil {
		err := fmt.Errorf(`pathChan is null`)
		logs.Error(err)
		errChan <- err
		return
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.Error(`panic:`, err)
				return
			}
		}()
		var path string
		for {
			select {
			case <-ctx.Done():
				return
			case p := <-pathChan:
				if path == p {
					continue
				}
				path = p
				logs.Info(`path:`, path)
				isExist, err := common.PathExists(path)
				if err != nil {
					logs.Error(err)
					errChan <- err
					return
				}
				if !isExist {
					err := fmt.Errorf(`path is not exist`)
					logs.Error(err)
					errChan <- err
					return
				}
				// 先打开这个文件，看看这个文本文件有多少个字节
				logFile, err := os.Open(path)
				if err != nil {
					logs.Error(err)
					logFile.Close()
					errChan <- err
					return
				}
				logFi, err := logFile.Stat()
				if err != nil {
					logs.Error(err)
					logFile.Close()
					errChan <- err
					return
				}
				// 设置返回的最大字节数
				var logFileSize int64
				confFileSize := common.ConfViper.GetInt64(`logMaxSize`)
				if confFileSize <= 0 {
					err := fmt.Errorf(`confFileSize is not right`)
					logs.Error(err)
					logFile.Close()
					errChan <- err
					return
				}
				if logFi.Size() >= confFileSize {
					logFileSize = confFileSize
				} else {
					logFileSize = logFi.Size()
				}
				logFile.Close()
				t, err := tail.TailFile(path, tail.Config{
					Follow:   true,
					ReOpen:   true,
					Location: &tail.SeekInfo{Offset: -logFileSize, Whence: 2},
				})
				if err != nil {
					logs.Error(err)
					errChan <- err
					return
				}
				tailChan <- t
				logs.Info(`tailChan is already`)
			}
		}
	}()
	return
}

func wsWrite(conn *websocket.Conn, tailChan chan *tail.Tail, ctx context.Context) (errChan chan error) {
	errChan = make(chan error)
	if conn == nil {
		err := fmt.Errorf(`conn is null`)
		logs.Error(err)
		errChan <- err
		return
	}
	if tailChan == nil {
		err := fmt.Errorf(`tailChan is null`)
		logs.Error(err)
		errChan <- err
		return
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.Error(err)
				return
			}
		}()
		var tail1 *tail.Tail
		// 给一个空的对象，避免报空指针错误
		tail1 = &tail.Tail{
			Filename: ``,
			Lines:    make(chan *tail.Line),
		}
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-tailChan:
				logs.Info(`tailChan is get`)
				if tail1.Filename != `` {
					logs.Info(tail1.Filename)
					err := tail1.Stop()
					logs.Info(`tail1 is stop`)
					if err != nil {
						logs.Error(err)
						errChan <- err
						return
					}
				}
				// logs.Info(`tail1 is stop`)
				tail1 = t
			case line := <-tail1.Lines:
				// logs.Info(`tail.Lines is get`)
				// logs.Info(line.Text)
				if err := conn.WriteJSON(line.Text); err != nil {
					logs.Error(err)
					errChan <- err
					return
				}
			}
		}
	}()
	return
}
