package lib

import (
	"net/http"
)

const (
	NamespaceUser = "user"
)

func GetUserCookie(r *http.Request) (sessId string) {
	cookie, _ := r.Cookie(NamespaceUser)
	if cookie == nil {
		return ""
	}

	sessId = cookie.Value
	return sessId
}

func SetUserCookie(w http.ResponseWriter, uuid string) {
	cookie := http.Cookie{
		Name:     NamespaceUser,
		Value:    uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func UnsetUserCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(NamespaceUser)
	if cookie != nil {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
}
