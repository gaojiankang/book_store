package controller

import (
	"goWeb/day08/bookStore/dao"
	"goWeb/day08/bookStore/lib"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var page *lib.BookPage
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}

	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetBooks(pageNo)
	} else {
		page, _ = dao.GetBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}

	page.IsLogin = dao.IsLogin(r)
	sessId := lib.GetUserCookie(r)
	page.Username, _ = dao.GetUsernameBySessionId(sessId)

	tpl := template.Must(template.ParseFiles("views/index.html"))
	tpl.Execute(w, page)
}
