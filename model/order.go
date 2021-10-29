package model

type Order struct {
	OrderId string
	CreateTime string
	TotalCount int64
	TotalAmount float64
	State int64
	UserName string
	UserId int
}
func (order *Order) NoSend() bool {
	return order.State == 0
}
func (order *Order) OkSend() bool {
	return order.State == 1
}
func (order *Order) Complete() bool {
	return order.State == 2
}

type OrderItem struct {
	OrderItemId int64
	Count int64
	Amount float64
	Title string
	Author string
	Price float64
	ImgPath string
	OrderId string
}
