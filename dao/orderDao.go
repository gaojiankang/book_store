package dao

import (
	"goWeb/day08/bookStore/model"
	"goWeb/day08/bookStore/utils"
)

func AddOrder(order *model.Order) (err error) {
	sql := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?, ?, ?, ?, ?, ?)"
	_, err = utils.DB.Exec(sql, order.OrderId, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserId)
	if err != nil {
		return
	}

	return nil
}

func GetOrders() ([]*model.Order, error) {
	sqlStr := "select id, total_count, total_amount, state, user_id, create_time from orders"
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderId, &order.TotalAmount, &order.TotalAmount, &order.State, &order.UserId, &order.CreateTime)
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderByUserId(userId int) ([]*model.Order, error) {
	sqlStr := "select id, total_count, total_amount, state, user_id, create_time from orders where user_id=?"
	rows, err := utils.DB.Query(sqlStr, userId)
	if err != nil {
		return nil, err
	}

	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderId, &order.TotalAmount, &order.TotalAmount, &order.State, &order.UserId, &order.CreateTime)
		orders = append(orders, order)
	}

	return orders, nil
}

func UpdateOrderState(orderId string, state int) (err error) {
	sqlStr := "update orders set state=? where id=?"
	_, err = utils.DB.Exec(sqlStr, state, orderId)
	if err != nil {
		return
	}
	return nil
}
