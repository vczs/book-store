package model

type Session struct {
	SessionId string
	UserName string
	UserId int //外键 和user的Id关联
}
