package dao

import (
	"goWeb/day08/bookStore/lib"
	"goWeb/day08/bookStore/model"
	"goWeb/day08/bookStore/utils"
	"net/http"
)

func AddSession(sess *model.Session) (err error) {
	sqlStr := "insert into sessions(session_id, username, user_id) values(?, ?, ?)"
	_, err = utils.DB.Exec(sqlStr, sess.SessionId, sess.Username, sess.UserId)
	if err != nil {
		return
	}
	return nil
}

func DeleteSession(sessId string) (err error) {
	sqlStr := "delete from sessions where session_id=?"
	_, err = utils.DB.Exec(sqlStr, sessId)
	if err != nil {
		return
	}
	return nil
}

func GetSession(sessId string) (sess *model.Session, err error) {
	sqlStr := "select session_id, username, user_id from sessions where session_id=?"
	sqlStmt, err := utils.DB.Prepare(sqlStr)
	if err != nil {
		return
	}
	row := sqlStmt.QueryRow(sessId)
	sess = &model.Session{}
	row.Scan(&sess.SessionId, &sess.Username, &sess.UserId)
	return sess, nil
}

func GetUsernameBySessionId(sessId string) (username string, err error) {
	sqlStr := "select username from sessions where session_id=?"
	row := utils.DB.QueryRow(sqlStr, sessId)
	err = row.Scan(&username)
	if err != nil {
		return
	}
	return username, nil
}

func IsLogin(r *http.Request) bool {
	sessId := lib.GetUserCookie(r)
	if sessId == "" {
		return false
	}
	session, _ := GetSession(sessId)
	if session.UserId < 0 {
		return false
	}
	return true
}
