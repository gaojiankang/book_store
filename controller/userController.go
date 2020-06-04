package controller

import (
	"goWeb/day08/bookStore/dao"
	"goWeb/day08/bookStore/lib"
	"goWeb/day08/bookStore/model"
	"goWeb/day08/bookStore/utils"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if ok := dao.IsLogin(r); ok {
		Index(w, r)
		return
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	var tpl *template.Template
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.Id > 0 {
		uuid := utils.CreateUUID()
		sess := &model.Session{
			SessionId: uuid,
			Username:  user.Username,
			UserId:    user.Id,
		}
		dao.AddSession(sess)

		lib.SetUserCookie(w, uuid)

		tpl = template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		tpl.Execute(w, user)
	} else {
		tpl = template.Must(template.ParseFiles("views/pages/user/login.html"))
		tpl.Execute(w, "用户名或密码不正确！")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	sessId := lib.GetUserCookie(r)
	dao.DeleteSession(sessId)
	lib.UnsetUserCookie(w, r)

	Index(w, r)
}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	var tpl *template.Template
	user, _ := dao.CheckUserName(username)
	if user.Id > 0 {
		tpl = template.Must(template.ParseFiles("views/pages/user/regist.html"))
		tpl.Execute(w, "username has been exists.")
	} else {
		dao.SaveUser(username, password, email)
		tpl = template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		tpl.Execute(w, "")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.Id > 0 {
		w.Write([]byte("the username has been exists."))
	} else {
		w.Write([]byte("<font style='color: green;'>the username can use.</font>"))
	}
}
