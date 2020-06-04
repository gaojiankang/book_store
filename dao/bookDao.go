package dao

import (
	"goWeb/day08/bookStore/lib"
	"goWeb/day08/bookStore/model"
	"goWeb/day08/bookStore/utils"
	"strconv"
)

const PageSize = 4

func GetBooks(sPageNo string) (page *lib.BookPage, err error) {
	var totalPage int
	var totalRecord int
	var books []*model.Book
	pageNo, _ := strconv.Atoi(sPageNo)

	sqlStr := "select count(*) from books"
	row := utils.DB.QueryRow(sqlStr)
	row.Scan(&totalRecord)

	if totalRecord%PageSize == 0 {
		totalPage = totalRecord / PageSize
	} else {
		totalPage = totalRecord/PageSize + 1
	}

	offset := (pageNo - 1) * PageSize

	sqlStr = "select id, title, author, price, sales, stock, img_path from books limit ?,?"
	rows, err := utils.DB.Query(sqlStr, offset, PageSize)
	if err != nil {
		return
	}

	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	page = &lib.BookPage{
		Books:       books,
		PageNo:      pageNo,
		PageSize:    PageSize,
		TotalPage:   totalPage,
		TotalRecord: totalRecord,
	}

	return page, nil
}

func GetBooksByPrice(sPageNo string, minPrice, maxPrice string) (page *lib.BookPage, err error) {
	var totalPage int
	var totalRecord int
	var books []*model.Book
	pageNo, _ := strconv.Atoi(sPageNo)

	sqlStr := "select count(*) from books where price between ? and ?"
	row := utils.DB.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)

	if totalRecord%PageSize == 0 {
		totalPage = totalRecord / PageSize
	} else {
		totalPage = totalRecord/PageSize + 1
	}

	offset := (pageNo - 1) * PageSize

	sqlStr = "select id, title, author, price, sales, stock, img_path from books where price between ? and ? limit ?,?"
	rows, err := utils.DB.Query(sqlStr, minPrice, maxPrice, offset, PageSize)
	if err != nil {
		return
	}

	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	page = &lib.BookPage{
		Books:       books,
		PageNo:      pageNo,
		PageSize:    PageSize,
		TotalPage:   totalPage,
		TotalRecord: totalRecord,
	}

	return page, nil
}

func AddBook(b *model.Book) (err error) {
	sqlStr := "insert into books(title, author, price, sales, stock, img_path) values(?, ?, ?, ?, ?, ?)"
	_, err = utils.DB.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return
	}
	return nil
}

func GetBookById(bookId string) (*model.Book, error) {
	sqlStr := "select id, title, author, price, sales, stock, img_path from books where id=?"
	row := utils.DB.QueryRow(sqlStr, bookId)
	book := &model.Book{}
	row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

func UpdateBook(b *model.Book) (err error) {
	sqlStr := "update books set title=?, author=?, price=?, sales=?, stock=? where id=?"
	_, err = utils.DB.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.Id)
	if err != nil {
		return
	}
	return nil
}

func DeleteBook(bookId string) (err error) {
	sqlStr := "delete from books where id=?"
	_, err = utils.DB.Exec(sqlStr, bookId)
	if err != nil {
		return
	}
	return nil
}
