package lib

import "goWeb/day08/bookStore/model"

type BookPage struct {
	Books       []*model.Book
	PageNo      int
	PageSize    int
	TotalPage   int
	TotalRecord int
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

func (p *BookPage) HasPrev() bool {
	return p.PageNo > 1
}

func (p *BookPage) HasNext() bool {
	return p.PageNo < p.TotalPage
}

func (p *BookPage) GetPrevPageNo() (prevPage int) {
	if p.HasPrev() {
		prevPage = p.PageNo - 1
	} else {
		prevPage = 1
	}
	return prevPage
}

func (p *BookPage) GetNextPageNo() (nextPage int) {
	if p.HasNext() {
		nextPage = p.PageNo + 1
	} else {
		nextPage = p.TotalPage
	}
	return nextPage
}
