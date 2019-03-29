package handler

import (
	"go-filestore-server/config"
	"go-filestore-server/db"
	"go-filestore-server/util"
	"net/http"
)

// 处理用户注册请求
func SignupHandler(w http.ResponseWriter,r *http.Request)  {
	if r.Method == http.MethodGet{
		/*
		data,err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		*/
		
		http.Redirect(w,r, "/static/view/signup.html",http.StatusFound)
		return
	}
	r.ParseForm()
	
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	
	if len(username) < 3 || len(passwd) < 5{
		w.Write([]byte("Invalid parameter"))
		return
	}
	
	// 对密码进行加盐及取sha1值加密
	encPasswd := util.Sha1([]byte(passwd+config.PwdSalt))
	// 将用户信息注册到用户表
	suc := db.UserSignup(username,encPasswd)
	if suc{
		w.Write([]byte("SUCCESS"))
	}else {
		w.Write([]byte("FAILED"))
	}
}

// 登陆接口
func SignInHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodGet{
		/*
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		*/
		http.Redirect(w,r,"/static/view/signin/html",http.StatusFound)
		return
	}
	
	r.ParseForm()
	
}