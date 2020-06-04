package dao

import (
	"goWeb/day08/bookStore/model"
	"goWeb/day08/bookStore/utils"
)

func CheckUserNameAndPassword(username, password string) (*model.User, error) {
	sqlStr := "select id, username, password, email from users where username=? and password=?"
	row := utils.DB.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func CheckUserName(username string) (*model.User, error) {
	sqlStr := "select id, username, password, email from users where username=?"
	row := utils.DB.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func SaveUser(username, password, email string) (err error) {
	sqlStr := "insert into users(username, password, email) values(?, ?, ?)"
	_, err = utils.DB.Exec(sqlStr, username, password, email)
	return
}
