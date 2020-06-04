package model

// Session 会话
type Session struct {
	SessionId string
	Username  string
	UserId    int
	Cart      *Cart
	OrderId   string
	Orders    []*Order
}
