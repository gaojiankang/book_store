package dao

import (
	"goWeb/day08/bookStore/model"
	"goWeb/day08/bookStore/utils"
)

func AddOrderItem(orderItem *model.OrderItem) (err error) {
	sqlStr := "insert into order_items(count, amount, title, author, price, img_path, order_id) values(?, ?, ?, ?, ?, ?, ?)"

	_, err = utils.DB.Exec(
		sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author,
		orderItem.Price, orderItem.ImgPath, orderItem.OrderId,
	)
	if err != nil {
		return
	}

	return nil
}

func GetOrderItemsByOrderId(orderId string) ([]*model.OrderItem, error) {
	sqlStr := "select id, count, amount, title, author, price, img_path, order_id from order_items where order_id=?"
	rows, err := utils.DB.Query(sqlStr, orderId)
	if err != nil {
		return nil, err
	}

	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(
			&orderItem.OrderItemId, &orderItem.Count, &orderItem.Amount, &orderItem.Title,
			&orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderId,
		)
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}
