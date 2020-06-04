package controller

import (
	"goWeb/day08/bookStore/dao"
	"goWeb/day08/bookStore/lib"
	"net/http"
	"strconv"
)

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	sCartItemId := r.FormValue("cartItemId")
	cartItemId, _ := strconv.Atoi(sCartItemId)
	sessId := lib.GetUserCookie(r)
	session, _ := dao.GetSession(sessId)
	userId := session.UserId
	cart, _ := dao.GetCartByUserId(userId)
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemId == cartItemId {
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems
			dao.DeleteCartItemById(cartItemId)
		}
	}

	dao.UpdateCartCount(cart)

	cart, _ = dao.GetCartByUserId(userId)
	if cart.TotalCount <= 0 {
		dao.DeleteCartByCartId(cart.CartId)
	}

	GetCartInfo(w, r)
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	sCartItemId := r.FormValue("cartItemId")
	cartItemId, _ := strconv.Atoi(sCartItemId)
	sBookCount := r.FormValue("bookCount")
	bookCount, _ := strconv.Atoi(sBookCount)

	sessId := lib.GetUserCookie(r)
	session, _ := dao.GetSession(sessId)
	userId := session.UserId
	cart, _ := dao.GetCartByUserId(userId)
	cartItems := cart.CartItems

	for _, v := range cartItems {
		if v.CartItemId == cartItemId {
			v.Count = bookCount

			dao.UpdateCartItemCount(v)
		}
	}

	dao.UpdateCartCount(cart)

	GetCartInfo(w, r)
}
