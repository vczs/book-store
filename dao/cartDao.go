package dao

import (
	"book-store/model"
	"book-store/utils"
)

//添加购物车
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values (?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cart.CartId, cart.GetCarTotalCount(), cart.GetCarTotalAmount(), cart.UserId)
	if err != nil {
		return err
	}
	cartItem := cart.CartItems //获取购物车中的所有购物项
	for _, v := range cartItem {
		AddCartItem(v) //将购物项插入到数据库中
	}
	return nil
}

//根据用户Id查询购物车
func GetCartByUserId(userId int) (*model.Cart, error) {
	//查询购物车
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, userId)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	if err != nil || cart.CartId == "" {
		return nil, err
	}
	//根据查询到的购物车id查询 属于该购物车id的所有购物项
	cartItems, err := GetCartItemsByCartId(cart.CartId)
	if err != nil {
		return nil, err
	}
	cart.CartItems = cartItems //将查询到的所有购物项赋值给cart.CartItems
	return cart, nil
}

//更新购物车的金额数量
func UpdateCart(cart *model.Cart) error {
	sql := "update carts set total_count = ?,total_amount= ? where id = ?"
	_, err := utils.Db.Exec(sql, cart.GetCarTotalCount(), cart.GetCarTotalAmount(), cart.CartId)
	if err != nil {
		return err
	}
	return nil
}

//根据购物车id删除购物车
func DeleteCartByCartId(cartId string) error {
	err := DeleteCartItemByCartId(cartId)
	if err != nil {
		return err
	}
	sql := "delete from carts where id = ?"
	_, err = utils.Db.Exec(sql, cartId)
	if err != nil {
		return err
	}
	return nil
}
