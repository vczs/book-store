package controller

import (
	"book-store/dao"
	"book-store/model"
	"book-store/utils"
	"html/template"
	"net/http"
	"os"
	"time"
)

//结账处理器
func Checkout(w http.ResponseWriter, r *http.Request) {
	session, _ := isLogin(r)
	//获取购物车
	cart, _ := dao.GetCartByUserId(session.UserId)
	if cart != nil {
		orderId := utils.CreateUUID()
		//创建订单
		order := &model.Order{
			OrderId:     orderId,
			CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
			TotalCount:  cart.TotalCount,
			TotalAmount: cart.TotalAmount,
			State:       0,
			UserId:      session.UserId,
			UserName:    session.UserName,
		}
		//保存订单到数据库
		dao.AddOrder(order)
		//创建订单项
		for _, v := range cart.CartItems {
			orderItem := &model.OrderItem{
				Count:   v.Count,
				Amount:  v.Amount,
				Title:   v.Book.Title,
				Author:  v.Book.Author,
				Price:   v.Book.Price,
				ImgPath: v.Book.ImgPath,
				OrderId: orderId,
			}
			//保存订单项数据库
			dao.AddOrderItem(orderItem)
			//更新图书库存和销量
			book := v.Book
			book.Sales = book.Sales + int(v.Count)
			book.Stock = book.Stock - int(v.Count)
			dao.UpdateBooks(book)
		}
		dao.DeleteCartByCartId(cart.CartId)
		//解析订单结算页面
		path, _ := os.Getwd()
		t := template.Must(template.ParseFiles(path + "/views/pages/cart/checkout.html"))
		t.Execute(w, order)
	}
}
