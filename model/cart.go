package model

type Cart struct {
	CartId string //购物车id
	CartItems []*CartItem //购物车中的购物项
	TotalCount int64 //购物车中书的总数
	TotalAmount float64 //购物车总金额
	UserId int //用户id 该购物车属于哪个用户
	UserName string //用户名

	IsEmpty bool
}
//获取购物车图书总金额
func (cart *Cart) GetCarTotalAmount() float64 {
	var totalAmount float64
	for _ , v := range cart.CartItems {
		totalAmount = totalAmount + v.GetCartItemAmount()
	}
	return totalAmount
}
//获取购物车图书总数量
func (cart *Cart) GetCarTotalCount() int64 {
	var totalCount int64
	for _ , v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}


//购物项
type CartItem struct {
	CartItemId int64 //购物项的Id
	Book *Book //购物项中的图书信息
	Count int64 //购物项中的图书数量
	Amount float64 //购物项中的金额小计
	CartId string //该购物项属于哪个购物车
}
//计算 购物项中的金额小计
func (cartItem *CartItem) GetCartItemAmount() float64 {
	return float64(cartItem.Count) * cartItem.Book.Price
}
