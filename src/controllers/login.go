package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../common"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/sessions"
)

// var store = sessions.NewCookieStore([]byte(sessionKey))
var store *sessions.CookieStore

func Login(w http.ResponseWriter, r *http.Request) {
	reply, err := NewReplyProto(`PUT`, `/allFile`)
	if err != nil {
		logs.Error(err)
		return
	}
	if r.Method != "PUT" {
		logs.Info(`somebody do not use GET method`)
		err = reply.ErrorResp(`you must use GET method`, w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	if body == nil || len(body) == 0 {
		err = fmt.Errorf(`body is null`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	var req ReqProto
	err = json.Unmarshal(body, &req)
	if err != nil {
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	ld, ok := req.Data.(map[string]interface{})
	if !ok {
		err = fmt.Errorf(`req.Data is not map[string]interface {} or do not exists`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	uid, ok := ld["uid"].(string)
	if !ok {
		err = fmt.Errorf(`uid is not string or do not exists`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	if uid == "" {
		err = fmt.Errorf(`uid is null`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	password, ok := ld["password"].(string)
	if !ok {
		err = fmt.Errorf(`password is not string or do not exists`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	if password == "" {
		err = fmt.Errorf(`password is null`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	ruid := common.ConfViper.GetString(`uid`)
	rpassword := common.ConfViper.GetString(`password`)
	if ruid == `` || rpassword == `` {
		err = fmt.Errorf(`you do not set your uid or password`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
	if ruid == uid && rpassword == password {
		session, err := store.Get(r, `isLogin`)
		// 设置过期时间
		session.Options = &sessions.Options{
			// Path:   "/login",
			MaxAge: sessionMaxAge, //以秒为单位
		}
		session.Values[`uid`] = uid
		session.Save(r, w)
		if err != nil {
			logs.Error(err)
			err = reply.ErrorResp(err.Error(), w)
			if err != nil {
				logs.Error(err)
				return
			}
			return
		}
		err = reply.SuccessResp(`登陆成功`, w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	} else {
		err = fmt.Errorf(`账号或者密码错误`)
		logs.Error(err)
		err = reply.ErrorResp(err.Error(), w)
		if err != nil {
			logs.Error(err)
			return
		}
		return
	}
}
