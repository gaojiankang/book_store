package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("test user dao func")
	t.Run("verification username and password:", testLogin)
	// t.Run("verification username:", testRegist)
	// t.Run("save user:", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123")
	fmt.Println("userinfo:", user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("userinfo:", user)
}

func testSave(t *testing.T) {
	err := SaveUser("admin", "123", "admin@gmail.com")
	fmt.Println(err)
}
