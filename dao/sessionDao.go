package dao

import (
	"book-store/model"
	"book-store/utils"
)

//添加session
func AddSession(session *model.Session) error {
	sqlStr := "insert into sessions value (?,?,?)"
	_, err := utils.Db.Exec(sqlStr, session.SessionId, session.UserName, session.UserId)
	if err != nil {
		return err
	}
	return nil
}

//删除session
func DeleteSession(sessionId string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessionId)
	if err != nil {
		return err
	}
	return nil
}

//根据id查询session
func GetSessionById(sessionId string) (*model.Session, error) {
	sqlStr := "select * from sessions where session_id = ?"
	row := utils.Db.QueryRow(sqlStr, sessionId)
	session := &model.Session{}
	err := row.Scan(&session.SessionId, &session.UserName, &session.UserId)
	if err != nil {
		return nil, err
	}
	return session, nil
}
