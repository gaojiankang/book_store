package controller

import (
	"goWeb/day08/bookStore/dao"
	"goWeb/day08/bookStore/lib"
	"goWeb/day08/bookStore/model"
	"net/http"
	"strconv"
	"text/template"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var page *lib.BookPage
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}

	page, _ = dao.GetBooks(pageNo)

	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	sPrice := r.PostFormValue("price")
	sSales := r.PostFormValue("sales")
	sStock := r.PostFormValue("stock")

	price, _ := strconv.ParseFloat(sPrice, 64)
	sales, _ := strconv.Atoi(sSales)
	stock, _ := strconv.Atoi(sStock)

	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   price,
		Sales:   sales,
		Stock:   stock,
		ImgPath: "/static/img/default.jpg",
	}
	dao.AddBook(book)

	GetBooks(w, r)
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	sBookId := r.FormValue("bookId")
	var book *model.Book
	book, _ = dao.GetBookById(sBookId)
	t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
	if book.Id <= 0 {
		book = nil
	}
	t.Execute(w, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	sBookId := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	sPrice := r.PostFormValue("price")
	sSales := r.PostFormValue("sales")
	sStock := r.PostFormValue("stock")

	bookId, _ := strconv.Atoi(sBookId)
	price, _ := strconv.ParseFloat(sPrice, 64)
	sales, _ := strconv.Atoi(sSales)
	stock, _ := strconv.Atoi(sStock)

	book := &model.Book{
		Id:      bookId,
		Title:   title,
		Author:  author,
		Price:   price,
		Sales:   sales,
		Stock:   stock,
		ImgPath: "/static/img/default.jpg",
	}
	dao.UpdateBook(book)

	GetBooks(w, r)
}

func SaveBook(w http.ResponseWriter, r *http.Request) {
	sBookId := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	sPrice := r.PostFormValue("price")
	sSales := r.PostFormValue("sales")
	sStock := r.PostFormValue("stock")

	bookId, _ := strconv.Atoi(sBookId)
	price, _ := strconv.ParseFloat(sPrice, 64)
	sales, _ := strconv.Atoi(sSales)
	stock, _ := strconv.Atoi(sStock)

	book := &model.Book{
		Id:      bookId,
		Title:   title,
		Author:  author,
		Price:   price,
		Sales:   sales,
		Stock:   stock,
		ImgPath: "/static/img/default.jpg",
	}
	if book.Id > 0 {
		dao.UpdateBook(book)
	} else {
		dao.AddBook(book)
	}

	GetBooks(w, r)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	sBookId := r.FormValue("bookId")

	dao.DeleteBook(sBookId)

	GetBooks(w, r)
}
