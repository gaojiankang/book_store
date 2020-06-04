package dao

import (
	"goWeb/day08/bookStore/model"
	"goWeb/day08/bookStore/utils"
)

func GetCartByUserId(userId int) (cart *model.Cart, err error) {
	sqlStr := "select id, total_count, total_amount, user_id from carts where user_id=?"
	row := utils.DB.QueryRow(sqlStr, userId)
	cart = &model.Cart{}
	err = row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	if err != nil {
		return nil, err
	}
	cartItems, _ := GetCartItemsByCartId(cart.CartId)
	cart.CartItems = cartItems
	return cart, nil
}

func AddCart(cart *model.Cart) (err error) {
	sqlStr := "insert into carts(id, total_count, total_amount, user_id) values(?, ?, ?, ?)"
	_, err = utils.DB.Exec(
		sqlStr, cart.CartId, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserId,
	)
	if err != nil {
		return
	}
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		AddCartItem(cartItem)
	}
	return nil
}

func UpdateCartCount(cart *model.Cart) (err error) {
	sqlStr := "update carts set total_count=?, total_amount=? where id=?"
	_, err = utils.DB.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartId)
	if err != nil {
		return
	}
	return nil
}

func DeleteCartByCartId(cartId string) (err error) {
	err = DeleteCartItemsByCartId(cartId)
	if err != nil {
		return
	}
	sqlStr := "delete from carts where id=?"
	_, err = utils.DB.Exec(sqlStr, cartId)
	if err != nil {
		return
	}
	return nil
}
