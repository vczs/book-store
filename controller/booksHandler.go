package controller

import (
	"book-store/dao"
	"book-store/model"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

////查询所有图书处理器
//func GetBooks(w http.ResponseWriter , r *http.Request) {
//	books , err := dao.GetBooks()
//	if err != nil {
//		fmt.Println("查询所有图书处理器出错:",err)
//		return
//	}
//	t := template.Must(template.ParseFiles("书城项目/views/pages/manager/book_manager.html"))
//	t.Execute(w,books)
//}
//查询带分页图书处理器
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取要显示的页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	} //如果pageNo为空就给个默认值1 即首页
	//调用GetPageBooks获取当前页的信息page
	page, err := dao.GetPageBooks(pageNo)
	if err != nil {
		fmt.Println("查询带分页图书处理器出错:", err)
		return
	}
	path, _ := os.Getwd()
	t := template.Must(template.ParseFiles(path + "/views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

//删除图书处理器
func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	err := dao.DeleteBooks(bookId)
	if err != nil {
		fmt.Println("删除图书处理器出错:", err)
		return
	}
	//刷新 调用查询处理器将更新后的数据库显示在客户端
	GetPageBooks(w, r)
}

//去更新或添加图书信息输入页处理器
func ToUpdateOrAddBooks(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	iBookId, _ := strconv.ParseInt(bookId, 10, 0)
	//判断是要更新 还是 添加
	if int(iBookId) > 0 { //客户端返回了Id  是更新操作
		book, err := dao.GetBooksById(bookId)
		if err != nil {
			fmt.Println("获取要更新的图书处理器出错:", err)
			return
		}
		path, _ := os.Getwd()
		t := template.Must(template.ParseFiles(path + "/views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else { //客户端没有返回Id  是添加操作
		path, _ := os.Getwd()
		t := template.Must(template.ParseFiles(path + "/views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
}

//更新或添加图书处理器
func UpdateOrAddBooks(w http.ResponseWriter, r *http.Request) {
	bookId := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	iBookId, _ := strconv.ParseInt(bookId, 10, 0)
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	books := &model.Book{Id: int(iBookId), Title: title, Author: author, Price: fPrice, Sales: int(iSales), Stock: int(iStock), ImgPath: "static/img/default.jpg"}
	if iBookId > 0 { //更新图书
		err := dao.UpdateBooks(books)
		if err != nil {
			fmt.Println("更新图书处理器出错:", err)
			return
		}
	} else { //添加图书
		//调用dao.AddBooks()向mysql中添加该书籍books
		err := dao.AddBooks(books)
		if err != nil {
			fmt.Println("添加图书处理器出错:", err)
			return
		}
	}
	//刷新 调用查询处理器将更新后的数据库显示在客户端
	GetPageBooks(w, r)
}

//获取图书价格范围分页处理器
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取要显示的页码
	pageNo := r.FormValue("pageNo")
	//获取价格范围
	min := r.FormValue("min")
	max := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	} //如果pageNo为空就给个默认值1 即首页
	var page *model.Page
	var err error
	if min == "" || max == "" {
		page, err = dao.GetPageBooks(pageNo)
	} else {
		page, err = dao.GetPageBooksByPrice(pageNo, min, max)
		if page != nil {
			page.MinPrice = min
			page.MaxPrice = max
		}
	}
	session, ok := isLogin(r)
	if ok && page != nil {
		//已经登录
		page.IsLogin = true
		page.UserName = session.UserName //将已登录账户的UserName放进page里 传给index.html
	}
	if err != nil {
		fmt.Println("获取图书价格范围分页处理器出错:", err)
	}
	path, _ := os.Getwd()
	t := template.Must(template.ParseFiles(path + "/views/pages/index.html"))
	t.Execute(w, page)
}
