package dao

import (
	"book-store/model"
	"book-store/utils"
)

//向数据库插入订单
func AddOrder(order *model.Order) error {
	sql := "insert into orders(id,create_time,total_count,total_amount,state,user_name,user_id) values (?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, order.OrderId, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserName, order.UserId)
	if err != nil {
		return err
	}
	return nil
}

//获取所有的订单
func GetOrders() ([]*model.Order, error) {
	sql := "select * from orders"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err = rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserName, &order.UserId)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

//根据userId获取订单
func GetOrderByUserId(userId int) ([]*model.Order, error) {
	sql := "select * from orders where user_id = ?"
	rows, err := utils.Db.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err = rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserName, &order.UserId)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

//根据订单id更新订单State
func UpdateOrderStateByOrderId(orderId string, state int64) error {
	sql := "update orders set state = ? where id = ?"
	_, err := utils.Db.Exec(sql, state, orderId)
	if err != nil {
		return err
	}
	return nil
}
