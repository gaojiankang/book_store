package model

type Book struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	Author  string  `json:"author"`
	Price   float64 `json:"price"`
	Sales   int     `json:"sales"`
	Stock   int     `json:"stock"`
	ImgPath string  `json:"img_path"`
}
