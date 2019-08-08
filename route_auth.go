package main

import (
	"go_web/data"
	"net/http"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	// UserByEmail 通过给定的电子邮件获取与之对应的 User 结构
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	// Encrypt 用于加密字符串
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, _ := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}

		// SetCookie 将 cookie 添加到响应的首部里面
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
