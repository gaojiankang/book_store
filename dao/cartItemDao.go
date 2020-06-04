package dao

import (
	"goWeb/day08/bookStore/model"
	"goWeb/day08/bookStore/utils"
)

func AddCartItem(cartItem *model.CartItem) (err error) {
	sqlStr := "insert into cart_items(count, amount, book_id, cart_id) values(?, ?, ?, ?)"
	_, err = utils.DB.Exec(
		sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId,
	)
	if err != nil {
		return
	}
	return nil
}

func GetCartItemByBookIdAndCartId(bookId string, cartId string) (cartItem *model.CartItem, err error) {
	sqlStr := "select id, count, amount, cart_id from cart_items where book_id=? and cart_id=?"
	row := utils.DB.QueryRow(sqlStr, bookId, cartId)
	cartItem = &model.CartItem{}
	err = row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartId)
	if err != nil {
		return nil, err
	}
	book, _ := GetBookById(bookId)
	cartItem.Book = book
	return cartItem, nil
}

func GetCartItemsByCartId(cartId string) (cartItems []*model.CartItem, err error) {
	sqlStr := "select id, count, amount, book_id, cart_id from cart_items where cart_id=?"
	rows, err := utils.DB.Query(sqlStr, cartId)
	if err != nil {
		return
	}

	for rows.Next() {
		var bookId string
		cartItem := &model.CartItem{}
		err = rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartId)
		if err != nil {
			return nil, err
		}
		book, _ := GetBookById(bookId)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func UpdateCartItemCount(cartItem *model.CartItem) (err error) {
	sqlStr := "update cart_items set count=?, amount=? where book_id=? and cart_id=?"
	_, err = utils.DB.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return
	}
	return nil
}

func DeleteCartItemsByCartId(cartId string) (err error) {
	sqlStr := "delete from cart_items where cart_id=?"
	_, err = utils.DB.Exec(sqlStr, cartId)
	if err != nil {
		return
	}
	return nil
}

func DeleteCartItemById(cartItemId int) (err error) {
	sqlStr := "delete from cart_items where id=?"
	_, err = utils.DB.Exec(sqlStr, cartItemId)
	if err != nil {
		return
	}
	return
}
