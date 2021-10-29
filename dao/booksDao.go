package dao

import (
	"book-store/model"
	"book-store/utils"
	"strconv"
)

//查询mysql中所有的书籍
func GetBooks() ([]*model.Book, error) {
	//编写sql语句
	sqlStr := "select * from books"
	//执行 查询多条
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建一个切片存储查询的图书
	var books []*model.Book
	//遍历查询结果
	for rows.Next() {
		book := &model.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

//添加书籍
func AddBooks(books *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) value (?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, books.Title, books.Author, books.Price, books.Sales, books.Stock, books.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//删除书籍
func DeleteBooks(booksId string) error {
	sqlStr := "delete from books where id = ?"
	_, err := utils.Db.Exec(sqlStr, booksId)
	if err != nil {
		return err
	}
	return nil
}

//通过id查询一本图书
func GetBooksById(booksId string) (*model.Book, error) {
	sqlStr := "select * from books where id = ?"
	row := utils.Db.QueryRow(sqlStr, booksId)
	book := &model.Book{}
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if err != nil {
		return nil, err
	}
	return book, nil
}

//更新书籍
func UpdateBooks(books *model.Book) error {
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id = ?"
	_, err := utils.Db.Exec(sqlStr, books.Title, books.Author, books.Price, books.Sales, books.Stock, books.Id)
	if err != nil {
		return err
	}
	return nil
}

//获取带分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//查询数据库表的总行数
	sqlStr := "select count(*) from books"
	//执行查询
	row := utils.Db.QueryRow(sqlStr)
	var totalRecord int64 //接收查询结果
	err := row.Scan(&totalRecord)
	if err != nil {
		return nil, err
	}
	var pageSize int64 = 4 //设置每页显示个数
	var totalPageNo int64  //接收总页数
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	if iPageNo > totalPageNo {
		iPageNo = totalPageNo
	} //如果要查询的页数大于总页数 就给他返回末页
	//查询该page的图书
	sqlStr2 := "select * from books limit ?,?"
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	//遍历查询到的该page得图书  并存储到model.Page的Books切片中
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	//model.Page的所有内容都获取到 现在可以创建实例化Page并返回了
	page := &model.Page{TotalRecord: totalRecord, TotalPageNo: totalPageNo, PageSize: pageSize, PageNo: iPageNo, Books: books}
	return page, nil
}

//获取图书价格范围分页
func GetPageBooksByPrice(pageNo string, min string, max string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//查询数据库表的总行数
	sqlStr := "select count(*) from books where price between ? and ?"
	//执行查询
	row := utils.Db.QueryRow(sqlStr, min, max)
	var totalRecord int64 //接收查询结果
	err := row.Scan(&totalRecord)
	if err != nil {
		return nil, err
	}
	var pageSize int64 = 4  //设置每页显示个数
	var totalPageNo int64   //接收总页数
	var books []*model.Book //接收查询到的书籍
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	if iPageNo > totalPageNo {
		iPageNo = totalPageNo
	} //如果要查询的页数大于总页数 就给他返回末页
	//查询该page的图书
	sqlStr2 := "select * from books where price between ? and ? limit ?,?"
	rows, err := utils.Db.Query(sqlStr2, min, max, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		//没查到符合条件的图书 返回初始值page  page不能返回nil 因为html无法解析会出错
		page := &model.Page{TotalRecord: 0, TotalPageNo: 0, PageSize: 0, PageNo: 0, Books: books}
		return page, err
	}
	//遍历查询到的该page得图书 并存储到model.Page的Books切片中
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	//model.Page的所有内容都获取到 现在可以创建实例化Page并返回了
	page := &model.Page{TotalRecord: totalRecord, TotalPageNo: totalPageNo, PageSize: pageSize, PageNo: iPageNo, Books: books}
	return page, nil
}
