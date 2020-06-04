package dao

import (
	"fmt"
	"goWeb/day08/bookStore/model"
	"testing"
)

func TestCartItem(t *testing.T) {
	fmt.Println("test cart item func")
	// t.Run("test add cart item", testAddCartItem)
	// t.Run("test update cart item count", testUpdateCartItemCount)
	// t.Run("test delete cart item", testDeleteCartItem)
}

func testAddCartItem(t *testing.T) {
	book := &model.Book{
		Id:    1,
		Price: 27.20,
	}
	cartItem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartId: "66667777",
	}
	AddCartItem(cartItem)
}

func testUpdateCartItemCount(t *testing.T) {
	book, _ := GetBookById("1")
	cartItem := &model.CartItem{
		Count:  15,
		Book:   book,
		CartId: "66667777",
	}
	UpdateCartItemCount(cartItem)
}

func testDeleteCartItem(t *testing.T) {
	DeleteCartItemsByCartId("66667777")
}
