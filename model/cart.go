package model

type Cart struct {
	CartId      string
	CartItems   []*CartItem
	TotalCount  int
	TotalAmount float64
	UserId      int
}

func (c *Cart) GetTotalCount() int {
	var totalCount int
	for _, v := range c.CartItems {
		totalCount += v.Count
	}
	return totalCount
}

func (c *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, cartItem := range c.CartItems {
		totalAmount += cartItem.GetAmount()
	}
	return totalAmount
}
