package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../common"

	"github.com/astaxie/beego/logs"
	"github.com/gorilla/sessions"
)

func ErrorHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logs.Error(err)
				var reply ReplyProto
				reply.Status = -1
				reply.Msg = fmt.Sprint(err)
				response, err := json.Marshal(reply)
				if err != nil {
					logs.Error(err)
					return
				}
				_, err = w.Write(response)
				if err != nil {
					logs.Error(err)
					return
				}
			}
		}()
		h(w, r)
	}
}

var sessionMaxAge int
var logMaxSize int64

func InitVariable() (err error) {
	sessionMaxAge = common.ConfViper.GetInt(`sessionMaxAge`)
	if sessionMaxAge <= 0 {
		err = fmt.Errorf(`sessionMaxAge must be greater than 0`)
		logs.Error(err)
		return err
	}
	sessionKey := common.ConfViper.GetString(`sessionKey`)
	if sessionKey == "" {
		err = fmt.Errorf(`sessionKey is null`)
		logs.Error(err)
		return err
	}
	logMaxSize = common.ConfViper.GetInt64(`logMaxSize`)
	if logMaxSize <= 0 {
		err = fmt.Errorf(`logMaxSize must be greater than 0`)
		logs.Error(err)
		return err
	}
	store = sessions.NewCookieStore([]byte(sessionKey))
	return nil
}
