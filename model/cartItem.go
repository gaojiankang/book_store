package model

type CartItem struct {
	CartItemId int
	Book       *Book
	Count      int
	Amount     float64
	CartId     string
}

func (ci *CartItem) GetAmount() float64 {
	price := ci.Book.Price
	return float64(ci.Count) * price
}
