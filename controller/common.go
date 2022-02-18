package controller

import (
	"book-store/dao"
	"book-store/model"
	"net/http"
)

// 判断客户端是否已经登录
func isLogin(r *http.Request) (*model.Session, bool) {
	//获取cookie 判断是否已登录
	cookie, err := r.Cookie("user")
	if err != nil {
		return nil, false
	}
	cookieValue := cookie.Value
	session, err := dao.GetSessionById(cookieValue)
	if err != nil {
		return nil, false
	}
	return session, true
}
