package controller

import (
	"goWeb/day08/bookStore/dao"
	"goWeb/day08/bookStore/lib"
	"goWeb/day08/bookStore/model"
	"goWeb/day08/bookStore/utils"
	"net/http"
	"text/template"
)

func AddBookToCart(w http.ResponseWriter, r *http.Request) {
	if ok := dao.IsLogin(r); !ok {
		w.Write([]byte("please login first."))
		return
	}

	bookId := r.FormValue("bookId")
	book, _ := dao.GetBookById(bookId)
	sessId := lib.GetUserCookie(r)
	sess, _ := dao.GetSession(sessId)
	userId := sess.UserId
	cart, _ := dao.GetCartByUserId(userId)
	if cart == nil {
		cartId := utils.CreateUUID()
		cart := &model.Cart{
			CartId: cartId,
			UserId: userId,
		}

		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book:   book,
			Count:  1,
			CartId: cartId,
		}
		cartItems = append(cartItems, cartItem)
		cart.CartItems = cartItems

		dao.AddCart(cart)
	} else {
		cartItem, _ := dao.GetCartItemByBookIdAndCartId(bookId, cart.CartId)
		if cartItem == nil {
			cartItem = &model.CartItem{
				Book:   book,
				Count:  1,
				CartId: cart.CartId,
			}
			cart.CartItems = append(cart.CartItems, cartItem)
			dao.AddCartItem(cartItem)
		} else {
			for _, v := range cart.CartItems {
				if v.Book.Id == cartItem.Book.Id {
					v.Count += 1
					dao.UpdateCartItemCount(v)
				}
			}
		}

		dao.UpdateCartCount(cart)
	}

	w.Write([]byte("You just added " + book.Title + " to cart"))
}

func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	sessId := lib.GetUserCookie(r)
	session, _ := dao.GetSession(sessId)

	var t *template.Template
	cart, _ := dao.GetCartByUserId(session.UserId)
	if cart != nil {
		session.Cart = cart
	}
	t = template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	t.Execute(w, session)
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")
	dao.DeleteCartByCartId(cartId)

	GetCartInfo(w, r)
}
