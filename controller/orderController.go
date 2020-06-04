package controller

import (
	"goWeb/day08/bookStore/dao"
	"goWeb/day08/bookStore/lib"
	"goWeb/day08/bookStore/model"
	"goWeb/day08/bookStore/utils"
	"net/http"
	"time"

	"github.com/alecthomas/template"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	sessId := lib.GetUserCookie(r)
	session, _ := dao.GetSession(sessId)
	userId := session.UserId
	cart, _ := dao.GetCartByUserId(userId)
	orderId := utils.CreateUUID()
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.Order{
		OrderId:     orderId,
		CreateTime:  timeStr,
		TotalAmount: cart.TotalAmount,
		TotalCount:  cart.TotalCount,
		State:       model.StateNotDeliver,
		UserId:      userId,
	}
	dao.AddOrder(order)

	cartItems := cart.CartItems
	for _, v := range cartItems {
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderId: orderId,
		}
		dao.AddOrderItem(orderItem)

		book := v.Book
		book.Sales += v.Count
		book.Stock -= v.Count
		dao.UpdateBook(book)
	}

	dao.DeleteCartByCartId(cart.CartId)
	session.OrderId = orderId
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, session)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, _ := dao.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

func GetMyOrder(w http.ResponseWriter, r *http.Request) {
	sessId := lib.GetUserCookie(r)
	session, _ := dao.GetSession(sessId)
	userId := session.UserId
	orders, _ := dao.GetOrderByUserId(userId)
	session.Orders = orders
	t := template.Must(template.ParseFiles("views/pages/order/my_order.html"))
	t.Execute(w, session)
}

func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	orderItems, _ := dao.GetOrderItemsByOrderId(orderId)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

func Deliver(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	dao.UpdateOrderState(orderId, model.StateDelivered)

	GetOrders(w, r)
}

func TakeDeliver(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	dao.UpdateOrderState(orderId, model.StateComplete)

	GetOrders(w, r)
}
