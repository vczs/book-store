package dao

import (
	"book-store/model"
	"book-store/utils"
)

//向数据库插入订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	sql := "insert into order_items(count,amount,title,author,price,img_path,order_id) values (?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderId)
	if err != nil {
		return err
	}
	return nil
}

//根据订单号获取该订单号所有的订单项
func GetOrderItemsByOrderId(orderId string) ([]*model.OrderItem, error) {
	sql := "select * from order_items where order_id = ?"
	rows, err := utils.Db.Query(sql, orderId)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		err = rows.Scan(&orderItem.OrderItemId, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderId)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
