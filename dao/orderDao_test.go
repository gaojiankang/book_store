package dao

import (
	"fmt"
	"goWeb/day08/bookStore/model"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	fmt.Println("test order func")
	// t.Run("test add order", testAddOrder)
	// t.Run("test get orders", testGetOrders)
	// t.Run("test get my order", testGetMyOrder)
	// t.Run("test get order items", testGetOrderItems)
	t.Run("test update order state", testUpdateOrderState)
}

func testAddOrder(t *testing.T) {
	orderId := "88888888"
	order := &model.Order{
		OrderId:     orderId,
		CreateTime:  time.Now().String(),
		TotalCount:  2,
		TotalAmount: 400,
		State:       model.StateDelivered,
		UserId:      2,
	}
	orderItem := &model.OrderItem{
		Count:   1,
		Amount:  300,
		Title:   "three kingdoms",
		Author:  "luo guan zhong",
		Price:   300,
		ImgPath: "/static/img/default.jpg",
		OrderId: orderId,
	}
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "journey to the west",
		Author:  "wu cheng en",
		Price:   100,
		ImgPath: "/static/img/default.jpg",
		OrderId: orderId,
	}

	AddOrder(order)
	AddOrderItem(orderItem)
	AddOrderItem(orderItem2)
}

func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, v := range orders {
		fmt.Println("order info:", v)
	}
}

func testGetOrderItems(t *testing.T) {
	orderItems, _ := GetOrderItemsByOrderId("88888888")
	for _, v := range orderItems {
		fmt.Println("order item info:", v)
	}
}

func testGetMyOrder(t *testing.T) {
	orders, _ := GetOrderByUserId(2)
	for _, v := range orders {
		fmt.Println("my order:", v)
	}
}

func testUpdateOrderState(t *testing.T) {
	UpdateOrderState("604de1cc-0a46-4703-7291-9049c5aa99e0", model.StateDelivered)
}
