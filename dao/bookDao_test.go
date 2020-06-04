package dao

import (
	"fmt"
	"goWeb/day08/bookStore/model"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println("test bookDao func")
	// t.Run("test get book with page", testGetBooks)
	// t.Run("test book with page and price", testGetPageBooksByPrice)
	// t.Run("test add book", testAddBook)
	// t.Run("test delete book", testDeleteBook)
	// t.Run("test get an book", testGetBook)
	// t.Run("test update book", testUpdateBook)
}

func testGetBooks(t *testing.T) {
	pageNo := "1"
	page, _ := GetBooks(pageNo)
	fmt.Println("current page no:", page.PageNo)
	fmt.Println("total page:", page.TotalPage)
	fmt.Println("total record:", page.TotalRecord)
	fmt.Println("current page book list")
	for _, v := range page.Books {
		fmt.Println("the book info:", v)
	}
}

func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetBooksByPrice("3", "10", "30")
	fmt.Println("current page no:", page.PageNo)
	fmt.Println("total page:", page.TotalPage)
	fmt.Println("total record:", page.TotalRecord)
	fmt.Println("current page book list")
	for _, v := range page.Books {
		fmt.Println("the book info:", v)
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "xi cheng",
		Author:  "carl",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	//调用添加图书的函数
	AddBook(book)
}

func testDeleteBook(t *testing.T) {
	DeleteBook("31")
}

func testGetPageBooks(t *testing.T) {
	book, _ := GetBookById("30")
	fmt.Println("get book info:", book)
}

func testUpdateBook(t *testing.T) {
	book := &model.Book{
		Id:      30,
		Title:   "教父2",
		Author:  "马里奥普佐",
		Price:   66.66,
		Sales:   10000,
		Stock:   1,
		ImgPath: "/static/img/default.jpg",
	}
	//调用更新图书的函数
	UpdateBook(book)
}
