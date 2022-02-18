package controller

import (
	"book-store/dao"
	"book-store/model"
	"book-store/utils"
	"crypto/md5"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//首页处理器 (与GetPageBooksByPrice()合并)
//func MainHandler(w http.ResponseWriter , r *http.Request){
//	//获取要显示的页码
//	pageNo := r.FormValue("pageNo")
//	if pageNo == "" {pageNo = "1"} //如果pageNo为空就给个默认值1 即首页
//	//调用GetPageBooks获取当前页的信息page
//	page , err := dao.GetPageBooks(pageNo)
//	if err != nil {
//		fmt.Println("查询带分页图书处理器出错:",err)
//		return
//	}
//	//解析模板
//	t := template.Must(template.ParseFiles("书城项目/views/pages/index.html"))
//	//执行
//	t.Execute(w,page)
//}

//登录处理器
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := isLogin(r) //判断是否登录
	if ok {
		//已经登录 就去首页
		GetPageBooksByPrice(w, r)
	} else {
		username := r.PostFormValue("username") //获取post表单输入的的username值
		password := r.PostFormValue("password") //获取post表单输入的的password值
		md5_pwd := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		user, err := dao.CheckUser(username, md5_pwd) //检查用户名密码
		if err == nil {
			//登陆成功
			uuid := utils.CreateUUID()
			//创建session
			session := &model.Session{
				SessionId: uuid,
				UserName:  user.UserName,
				UserId:    user.Id,
			}
			dao.AddSession(session) //将session存入mysql
			//创建一个cookie关联session
			cookie := &http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//将cookie发送给客户端
			http.SetCookie(w, cookie)
			path, _ := os.Getwd()
			t := template.Must(template.ParseFiles(path + "/views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			//登陆失败
			path, _ := os.Getwd()
			t := template.Must(template.ParseFiles(path + "/views/pages/user/login.html"))
			t.Execute(w, "用户名或密码不正确！")
		}
	}
}

//注册处理器
func RegistHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username") //获取post表单输入的的username值
	password := r.PostFormValue("password") //获取post表单输入的的password值
	email := r.PostFormValue("email")       //获取post表单输入的的email值
	if !dao.CheckUserName(username) {
		//用户不存在 可以注册
		md5_pwd := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		user := &model.User{UserName: username, Password: md5_pwd, Email: email}
		err := dao.SaveUser(user) //将用户存储到mysql
		if err == nil {
			//存储成功 表示注册成功
			path, _ := os.Getwd()
			t := template.Must(template.ParseFiles(path + "/views/pages/user/regist_success.html"))
			t.Execute(w, user)
			return
		}
	}
	//用户已经存在不能注册 或 存储失败 都视为注册失败
	path, _ := os.Getwd()
	t := template.Must(template.ParseFiles(path + "/views/pages/user/regist.html"))
	t.Execute(w, "该用户已存在！")
}

//注销处理器
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value的uuid值 并在mysql找到删除它
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
		//让客户端cookie失效
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	GetPageBooksByPrice(w, r)
}

//Ajax验证用户名是否可用处理器
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username") //获取post表单输入的的username值
	if dao.CheckUserName(username) {
		w.Write([]byte("用户名已存在！"))
	} else {
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}
