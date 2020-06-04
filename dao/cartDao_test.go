package dao

import (
	"fmt"
	"goWeb/day08/bookStore/model"
	"testing"
)

var CartId string = "66668888"

func TestCart(t *testing.T) {
	fmt.Println("测试购物车的相关函数")
	// t.Run("test add cart", testAddCart)
	// t.Run("test get cart item by book id", testGetCartItemByBookIdAndCartId)
	// t.Run("test get all items by cart id", testGetCartItemsByCartId)
	// t.Run("test get cart by user id", testGetCartByUserId)
	t.Run("test update cart item info by book id", testDeleteCartByCartId)
}

func testAddCart(t *testing.T) {
	t.Helper()
	book := &model.Book{
		Id:    1,
		Price: 27.20,
	}
	book2 := &model.Book{
		Id:    2,
		Price: 23.00,
	}

	var cartItems []*model.CartItem
	cartItem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartId: CartId,
	}
	cartItems = append(cartItems, cartItem)

	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartId: CartId,
	}
	cartItems = append(cartItems, cartItem2)

	cart := &model.Cart{
		CartId:    CartId,
		CartItems: cartItems,
		UserId:    1,
	}

	AddCart(cart)
}

func testGetCartItemByBookIdAndCartId(t *testing.T) {
	cartItem, _ := GetCartItemByBookIdAndCartId("1", CartId)
	fmt.Println("book id 1 cart items:", cartItem)
}

func testGetCartItemsByCartId(t *testing.T) {
	cartItems, _ := GetCartItemsByCartId(CartId)
	for _, v := range cartItems {
		fmt.Printf("cart item: %v\n", v)
	}
}

func testGetCartByUserId(t *testing.T) {
	cart, _ := GetCartByUserId(2)
	fmt.Printf("user id: %d cart info: %v\n", 2, cart)
}

func testUpdateCartCount(t *testing.T) {

}

func testDeleteCartByCartId(t *testing.T) {
	cartId := "0d1d2626-faf3-4c39-772f-438ec0a691b9"
	DeleteCartByCartId(cartId)
}
