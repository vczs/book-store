package dao

import (
	"book-store/model"
	"book-store/utils"
)

//向cart_items添加购物项
func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values (?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetCartItemAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

//根据书的id查购物项
func GetCartItemByBookId(bookId string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ?"
	row := utils.Db.QueryRow(sqlStr, bookId)
	carItem := &model.CartItem{}
	err := row.Scan(&carItem.CartItemId, &carItem.Count, &carItem.Amount, &carItem.CartId)
	if err != nil {
		return nil, err
	}
	return carItem, nil
}

//根据购物车id获取所有购物项
func GetCartItemsByCartId(cartId string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	rows, err := utils.Db.Query(sqlStr, cartId)
	if err != nil {
		return nil, err
	}
	var carItems []*model.CartItem
	var bookId string //接收查询购物项的图书的id
	for rows.Next() {
		carItem := &model.CartItem{}
		err := rows.Scan(&carItem.CartItemId, &carItem.Count, &carItem.Amount, &bookId, &carItem.CartId)
		if err != nil {
			return nil, err
		}
		book, err := GetBooksById(bookId) //根据查询到的购物项的图书的id获取图书信息
		if err != nil {
			return nil, err
		}
		carItem.Book = book //将图书信息赋值给carItem.Book
		carItems = append(carItems, carItem)
	}
	return carItems, nil
}

//根据书的id和购物车id查购物项
func GetCartItemByBookIdAndCartId(bookId string, cartId string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	row := utils.Db.QueryRow(sqlStr, bookId, cartId)
	carItem := &model.CartItem{}
	err := row.Scan(&carItem.CartItemId, &carItem.Count, &carItem.Amount, &carItem.CartId)
	if err != nil {
		return nil, err
	}
	book, _ := GetBooksById(bookId)
	carItem.Book = book
	return carItem, nil
}

//更新购物项的数量 根据书的id和购物车id
func UpdateCartItem(cartItem *model.CartItem) error {
	sql := "update cart_items set count = ? , amount = ? where book_id = ? and cart_id = ?"
	_, err := utils.Db.Exec(sql, cartItem.Count, cartItem.GetCartItemAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

//根据购物车id删除购物项
func DeleteCartItemByCartId(cartId string) error {
	sql := "delete from cart_items where cart_id = ?"
	_, err := utils.Db.Exec(sql, cartId)
	if err != nil {
		return err
	}
	return nil
}

//根据购物项id删除购物项
func DeleteCartItemByCartItemId(cartItemId string) error {
	sql := "delete from cart_items where id = ?"
	_, err := utils.Db.Exec(sql, cartItemId)
	if err != nil {
		return err
	}
	return nil
}
