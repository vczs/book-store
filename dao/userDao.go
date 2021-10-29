package dao

import (
	"book-store/model"
	"book-store/utils"
	"errors"
)

//检查用户名是否存在
func CheckUserName(userName string) bool {
	sqlStr := "select username from users where username = ?"
	row := utils.Db.QueryRow(sqlStr, userName)
	var temp string
	row.Scan(&temp)
	if temp == userName {
		return true
	}
	return false
}

//登录 先检查用户名是否存在 如果存在再检查密码是否正确
func CheckUser(userName string, password string) (*model.User, error) {
	if !CheckUserName(userName) {
		return nil, errors.New("用户不存在！\n")
	}
	//如果该用户存在 再检查密码是否正确
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, userName, password)
	user := &model.User{}
	err := row.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	if err != nil {
		return nil, errors.New("检查密码是否正确出错！\n")
	}
	return user, nil
}

//注册 将user存储到mysql
func SaveUser(user *model.User) error {
	sqlStr := "insert into users(username,password,email) value (?,?,?)"
	_, err := utils.Db.Exec(sqlStr, user.UserName, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}
