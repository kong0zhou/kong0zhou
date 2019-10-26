package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"
)

type SseData struct {
	Event string
	ID    int64
	Retry uint
	Data  string
}

var sseText string = `
id:%d
event:%s
retry:%d
data:%s

`

func (s SseData) convertText() (text string, err error) {
	if s.Data == `` {
		err = fmt.Errorf(`sseData is null`)
		logs.Error(err)
		return ``, err
	}
	text = fmt.Sprintf(sseText, s.ID, s.Event, s.Retry, s.Data)
	// logs.Info(text)
	return text, nil
}

type Sse struct {
	IsClosed  <-chan bool
	f         http.Flusher
	w         http.ResponseWriter
	touchTime time.Time
}

func NewSse(w http.ResponseWriter) (s *Sse, err error) {
	if w == nil {
		err = fmt.Errorf(`w is null`)
		logs.Error(err)
		return nil, err
	}
	fl, ok := w.(http.Flusher)
	if !ok {
		err = fmt.Errorf(`sse is unsupported`)
		logs.Error(err)
		return nil, err
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	w.Header().Set("Transfer-Encoding", "chunked")

	// 必须的执行这一步，sse链接是从服务端发送第一条信息之后才开始建立的
	fmt.Fprintf(w, "event:isconnected\ndata: connection is established\n\n")
	fl.Flush()

	isclosed := make(chan bool, 2)

	// 将一个 sse close的信号量变成两个，一个用于停止sse心跳包
	go func() {
		<-w.(http.CloseNotifier).CloseNotify()
		isclosed <- true
		isclosed <- true
	}()

	sse := &Sse{
		// IsClosed: w.(http.CloseNotifier).CloseNotify(),
		IsClosed:  isclosed,
		f:         fl,
		w:         w,
		touchTime: time.Now(),
	}

	//心跳包
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for {
			select {
			case <-sse.IsClosed:
				logs.Info(`sse通道已经关闭，心跳功能停止`)
				return
			case <-ticker.C:
				now := time.Now()
				nop := "event:nop\ndata:nop\n\n"
				if now.Sub(sse.touchTime) < 30*time.Second {
					continue
				}
				logs.Info("push keep alive msg")
				fmt.Fprint(w, nop)
				fl.Flush()
			}
		}
	}()

	return sse, nil
}

func (s *Sse) Write(data SseData) (err error) {
	if data.Data == `` {
		logs.Warn(`see: data is null or empty`)
	}
	select {
	case <-s.IsClosed:
		err = fmt.Errorf(`sse is closed`)
		logs.Error(err)
		return err
	default:
		sseText, err := data.convertText()
		if err != nil {
			logs.Error(err)
			return err
		}
		_, err = fmt.Fprint(s.w, sseText)
		if err != nil {
			logs.Error(err)
			return err
		}
		s.f.Flush()
		return nil
	}
}
