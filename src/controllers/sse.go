package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/logs"
)

type SseData struct {
	Event string
	ID    string
	Retry uint
	Data  string
}

var sseText string = `
id:%s

event:%s

Retry:%d

data:%s
`

func (s SseData) convertText() (text string, err error) {
	if s.Data == `` {
		err = fmt.Errorf(`sseData is null`)
		logs.Error(err)
		return ``, err
	}
	text = fmt.Sprintf(sseText, s.ID, s.Event, s.Retry, s.Data)
	logs.Info(text)
	return text, nil
}

type Sse struct {
	IsClosed <-chan bool
	f        http.Flusher
	w        http.ResponseWriter
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
	return &Sse{
		IsClosed: w.(http.CloseNotifier).CloseNotify(),
		f:        fl,
		w:        w,
	}, nil
}

func (s *Sse) Write(data SseData) (err error) {
	if data.Data == `` {
		err = fmt.Errorf(`see: data is null or empty`)
		logs.Error(err)
		return err
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
		fmt.Fprint(s.w, sseText)
		s.f.Flush()
		return nil
	}
}
