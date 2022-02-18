package controller

import (
	"book-store/dao"
	"book-store/model"
	"html/template"
	"net/http"
	"os"
)

type MyOrder struct {
	UserName string
	Orders   []*model.Order
	IsEmpty  bool
}

type Order struct {
	Orders  []*model.Order
	IsEmpty bool
}

//订单管理处理器
func GetOrder(w http.ResponseWriter, r *http.Request) {
	orders, err := dao.GetOrders()
	order := &Order{}
	if len(orders) < 1 || err != nil {
		order.IsEmpty = true
	} else {
		order.IsEmpty = false
		order.Orders = orders
	}
	path, _ := os.Getwd()
	t := template.Must(template.ParseFiles(path + "/views/pages/order/order_manager.html"))
	t.Execute(w, order)
}

//订单详情处理器
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	orderItems, _ := dao.GetOrderItemsByOrderId(orderId)
	if orderItems != nil {
		path, _ := os.Getwd()
		t := template.Must(template.ParseFiles(path + "/views/pages/order/order_info.html"))
		t.Execute(w, orderItems)
	}
}

//我的订单处理器
func GetMyOrder(w http.ResponseWriter, r *http.Request) {
	myOrder := &MyOrder{}
	session, _ := isLogin(r)
	if session != nil {
		userId := session.UserId
		orders, getOrdersErr := dao.GetOrderByUserId(userId)
		if getOrdersErr != nil || orders == nil {
			myOrder.UserName = session.UserName
			myOrder.IsEmpty = true
		}
		myOrder.Orders = orders
		path, _ := os.Getwd()
		t := template.Must(template.ParseFiles(path + "/views/pages/order/order.html"))
		t.Execute(w, myOrder)
	} else {
		path, _ := os.Getwd()
		t := template.Must(template.ParseFiles(path + "/views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

//发货处理器
func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	dao.UpdateOrderStateByOrderId(orderId, 1)
	GetOrder(w, r)
}

//确认收货处理器
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	dao.UpdateOrderStateByOrderId(orderId, 2)
	GetMyOrder(w, r)
}
