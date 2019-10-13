package controllers

import (
	"fmt"
	"net/http"

	"../common"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/sessions"
)

func SessionCheck(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reply, err := NewReplyProto(`no`, `no`)
		if err != nil {
			logs.Error(err)
			return
		}
		session, err := store.Get(r, `isLogin`)
		if err != nil {
			logs.Error(err)
			err = reply.ErrorResp(err.Error(), w)
			if err != nil {
				logs.Error(err)
				return
			}
			return
		}
		uid, ok := session.Values[`uid`].(string)
		if !ok {
			logs.Info(`重新登陆`)
			logs.Info(session.Values[`uid`])
			err = reply.ErrorResp(`请重新登陆`, w)
			if err != nil {
				logs.Error(err)
				return
			}
			return
		}
		ruid := common.ConfViper.GetString(`uid`)
		if ruid == `` {
			err = fmt.Errorf(`ruid is null`)
			logs.Error(err)
			err = reply.ErrorResp(err.Error(), w)
			if err != nil {
				logs.Error(err)
				return
			}
			return
		}
		if uid != common.ConfViper.GetString(`uid`) {
			logs.Info(`重新登陆`)
			err = reply.ErrorResp(`请重新登陆`, w)
			if err != nil {
				logs.Error(err)
				return
			}
			return
		}
		session.Options = &sessions.Options{
			MaxAge: sessionMaxAge, //以秒为单位
		}
		session.Save(r, w)
		h(w, r)
	}
}

func CheckUser(w http.ResponseWriter, r *http.Request) {
	reply, err := NewReplyProto(`no`, `no`)
	if err != nil {
		logs.Error(err)
		return
	}
	session, err := store.Get(r, `isLogin`)
	if err != nil {
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	uid, ok := session.Values[`uid`].(string)
	if !ok {
		logs.Info(`重新登陆`)
		logs.Info(session.Values[`uid`])
		err = reply.ErrorResp(`请重新登陆`, w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	ruid := common.ConfViper.GetString(`uid`)
	if ruid == `` {
		err = fmt.Errorf(`ruid is null`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	if uid != common.ConfViper.GetString(`uid`) {
		logs.Info(`重新登陆`)
		err = reply.ErrorResp(`请重新登陆`, w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	err = reply.SuccessResp(nil, w)
	if err != nil {
		logs.Error(err)
		return
	}
	return
}
