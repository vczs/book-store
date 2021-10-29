package controller

import (
	"book-store/dao"
	"book-store/model"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Data struct {
	TotalCount  int64
	TotalAmount float64
	Amount      float64
}

//添加图书到购物车处理器
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//获取session得到当前用户id
	session, ok := isLogin(r)
	if session != nil && ok {
		//用户已登录
		//获取要添加图书的id
		bookId := r.PostFormValue("bookId")
		//根据图书id获取图书信息
		book, _ := dao.GetBooksById(bookId)
		//判断mysql中是否有该用户购物车
		cart, _ := dao.GetCartByUserId(session.UserId) //根据用户id获取用户购物车
		if cart != nil {                               //当前用户有购物
			//判断该用户购物车中是否有当前图书购物项
			getCartItem, _ := dao.GetCartItemByBookIdAndCartId(bookId, cart.CartId)
			if getCartItem != nil {
				//有该购物项  数量加1
				cts := cart.CartItems
				for _, v := range cts {
					if v.Book.Id == getCartItem.Book.Id {
						//当前购物项的数量加1
						v.Count = v.Count + 1
						dao.UpdateCartItem(v)
					}
				}
			} else {
				//没有该购物项  创建购物项
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartId: cart.CartId,
				}
				cart.CartItems = append(cart.CartItems, cartItem)
				dao.AddCartItem(cartItem)
			}
			dao.UpdateCart(cart)
		} else {
			//当前用户没有购物车 创建购物车
			cart := &model.Cart{
				CartId: session.SessionId,
				UserId: session.UserId,
			}
			//创建购物项 将购物项添加到*model.CartItem类型的切片中
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartId: cart.CartId,
			}
			cartItems = append(cartItems, cartItem)
			//将cartItems切片赋值给cart的CartItems字段
			cart.CartItems = cartItems
			dao.AddCart(cart)
		}
		w.Write([]byte("您刚刚将 " + book.Title + " 加入到购物车"))
	} else {
		//用户未登录
		w.Write([]byte("未登录"))
	}
}

//获取购物车信息处理器
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	session, ok := isLogin(r)
	if ok {
		//获取购物车
		cart, _ := dao.GetCartByUserId(session.UserId)
		if cart == nil {
			cart = &model.Cart{
				UserId:  session.UserId,
				IsEmpty: true,
			}
		}
		if cart.CartItems == nil {
			cart.IsEmpty = true
		}
		cart.UserName = session.UserName
		path, _ := os.Getwd()
		t := template.Must(template.ParseFiles(path + "/views/pages/cart/cart.html"))
		t.Execute(w, cart)
	} else {
		path, _ := os.Getwd()
		t := template.Must(template.ParseFiles(path + "/views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

//清空购物车处理器
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")
	dao.DeleteCartByCartId(cartId)
	GetCartInfo(w, r)
}

//删除购物项处理器
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")
	iCartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)
	session, _ := isLogin(r)
	//获取session里该用户的购物车
	userId := session.UserId
	cart, _ := dao.GetCartByUserId(userId)
	if cart != nil {
		cartItems := cart.CartItems
		//遍历购物车的购物项找到要删除的购物项 在购物车的购物项中删除它
		for k, v := range cartItems {
			if v.CartItemId == iCartItemId {
				cartItems = append(cartItems[:k], cartItems[k+1:]...)
				//从数据库的购物项中删除该购物项
				dao.DeleteCartItemByCartItemId(cartItemId)
			}
		}
		cart.CartItems = cartItems
	}
	dao.UpdateCart(cart)
	GetCartInfo(w, r)
}

//更新购物车处理器
func UpdateCart(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")
	bookCount := r.FormValue("bookCount")
	iCartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)
	iCookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	session, _ := isLogin(r)
	//获取session里该用户的购物车
	userId := session.UserId
	cart, _ := dao.GetCartByUserId(userId)
	if cart != nil {
		cartItems := cart.CartItems
		//遍历购物车的购物项找到要更新的购物项 在购物车的购物项中更新它
		for _, v := range cartItems {
			if v.CartItemId == iCartItemId {
				//更新目标购物项
				v.Count = iCookCount
				dao.UpdateCartItem(v)
			}
		}
	}
	dao.UpdateCart(cart)
	//将更新后的信息传给客户端
	tempCart, _ := dao.GetCartByUserId(userId)
	var amount float64
	tempCartItem := tempCart.CartItems
	for _, v := range tempCartItem {
		if v.CartItemId == iCartItemId {
			amount = v.Amount
		}
	}
	data := Data{
		TotalCount:  tempCart.TotalCount,
		TotalAmount: tempCart.TotalAmount,
		Amount:      amount,
	}
	js, _ := json.Marshal(data)
	w.Write(js)
}
