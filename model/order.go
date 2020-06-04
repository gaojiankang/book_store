package model

const (
	// StateNotDeliver 0 未发货
	StateNotDeliver = 0
	// StateDelivered 1 已发货
	StateDelivered = 1
	// StateComplete 2 交易完成
	StateComplete = 2
)

//Order 结构
type Order struct {
	OrderId     string  //订单号
	CreateTime  string  //生成订单的时间
	TotalCount  int     //订单中的总数量
	TotalAmount float64 //订单中的总金额
	State       int     //订单的状态 0 未发货 1 已发货 2 交易完成
	UserId      int     //订单所属的用户
}

func (order *Order) NotDeliver() bool {
	return order.State == StateNotDeliver
}

func (order *Order) Delivered() bool {
	return order.State == StateDelivered
}

func (order *Order) Complete() bool {
	return order.State == StateComplete
}
