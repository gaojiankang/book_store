package main

import (
	"goWeb/day08/bookStore/controller"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/", controller.Index)

	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/regist", controller.Regist)
	http.HandleFunc("/checkUserName", controller.CheckUserName)

	http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/addBook", controller.AddBook)
	http.HandleFunc("/editBook", controller.EditBook)
	http.HandleFunc("/saveBook", controller.SaveBook)
	http.HandleFunc("/deleteBook", controller.DeleteBook)

	http.HandleFunc("/addBookToCart", controller.AddBookToCart)
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	http.HandleFunc("/deleteCart", controller.DeleteCart)

	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)

	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getOrders", controller.GetOrders)
	http.HandleFunc("/getMyOrder", controller.GetMyOrder)
	http.HandleFunc("/deliver", controller.Deliver)
	http.HandleFunc("/takeDeliver", controller.TakeDeliver)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)

	http.ListenAndServe(":8080", nil)
}
