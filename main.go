package main

import (
	"book-store/controller"
	"net/http"
	"os"
)

func main() {
	parseStatic()
	userHandler()
	bookHandler()
	cartHandler()
	orderHandler()
	//8080端口监听 多路复用器为nil表示使用默认值
	http.ListenAndServe(":8080", nil)
}

func parseStatic() {
	//设置处理静态资源(css和js文件)
	//http.StripPrefix(prefix string , h Handler) Handler
	//http.StripPrefix返回一个处理器，该处理器会将请求的URL.Path字段中给定前缀prefix去除后，再交给h处理。如果prefix为空会回复404。
	//此时html中的所有URL地址以static开头的资源都能被加载出来(pages同理)
	//逻辑：如果访问"/static/"开头的资源，就去"书城项目/views/static"里找(pages同理)
	path, _ := os.Getwd()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path+"/views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir(path+"/views/pages"))))
}

func userHandler() {
	//去首页
	http.HandleFunc("/", controller.GetPageBooksByPrice)
	//去登录
	http.HandleFunc("/login", controller.LoginHandler)
	//去注册
	http.HandleFunc("/regist", controller.RegistHandler)
	//去注销
	http.HandleFunc("/logout", controller.LogoutHandler)
	//Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
}

func bookHandler() {
	////去图书管理(所有的)
	//http.HandleFunc("/getPageBooks",controller.GetBooks)
	//去图书管理(带分页的)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//去删除图书
	http.HandleFunc("/deleteBooks", controller.DeleteBooks)
	//去更新或添加图书信息输入页
	http.HandleFunc("/toUpdateOrAddBooks", controller.ToUpdateOrAddBooks)
	//去更新或添加图书
	http.HandleFunc("/updateOrAddBooks", controller.UpdateOrAddBooks)
	//去获取图书价格范围分页
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
}

func cartHandler() {
	//去添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	//去获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	//去清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	//去删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	//更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCart)
}

func orderHandler() {
	//去结账
	http.HandleFunc("/checkout", controller.Checkout)
	//去订单管理
	http.HandleFunc("/getOrder", controller.GetOrder)
	//去获取订单详情
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	//去我的订单
	http.HandleFunc("/getMyOrder", controller.GetMyOrder)
	//去发货
	http.HandleFunc("/sendOrder", controller.SendOrder)
	//去确认收货
	http.HandleFunc("/takeOrder", controller.TakeOrder)
}
