package model

type Page struct {
	TotalRecord int64 //总数 通过查询mysql得到
	TotalPageNo int64 //总页数 通过计算总数得到
	PageSize int64 //每页要显示的个数
	PageNo int64 //当前页
	Books []*Book //当前页要显示的图书信息
	MinPrice string //最低价格查询
	MaxPrice string //最高价格查询

	IsLogin bool //是否登录
	UserName string //已登录账户的用户名
}

//是否有上一页
func (page *Page) IsHasPrev() bool {
	return page.PageNo > 1
}
//是否有下一页
func (page *Page) IsHasNext() bool {
	return page.PageNo < page.TotalPageNo
}
//获取上一页
func (page *Page) GetPrev() int64 {
	if !page.IsHasPrev(){return page.PageNo}
	return page.PageNo - 1
}
//获取下一页
func (page *Page) GetNext() int64 {
	if !page.IsHasNext(){return page.PageNo}
	return page.PageNo + 1
}